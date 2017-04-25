package xnet

import (
	"errors"
	"libs/xlog"
	"net"
	"sync/atomic"
)

var (
	allocSessionID uint64
	// ErrSessionClosed TODO
	ErrSessionClosed = errors.New("session closed")
	ErrClosedLock    = errors.New("session close lock")
)

// Session TODO
type Session struct {
	id        uint64
	codec     Codec
	conn      net.Conn
	sendChan  chan *Packet
	closeChan chan int
	closed    int32
	mgr       *SessionMgr
}

func newSession(mgr *SessionMgr, conn net.Conn, codec Codec, sndChanSize int) *Session {
	c := &Session{
		codec:     codec,
		mgr:       mgr,
		id:        atomic.AddUint64(&allocSessionID, 1),
		closeChan: make(chan int),
		closed:    0,
		conn:      conn,
	}

	if sndChanSize > 0 {
		c.sendChan = make(chan *Packet, sndChanSize)
		go c.sender()
	}

	return c
}

// ID TODO
func (c *Session) ID() uint64 {
	return c.id
}

// RemoteAddr TODO
func (c *Session) RemoteAddr() string {
	return c.conn.RemoteAddr().String()
}

// Receive TODO
func (c *Session) Receive() (*Packet, error) {
	return c.codec.Receive()
}

// Close TODO
func (c *Session) Close() error {
	if atomic.CompareAndSwapInt32(&c.closed, 0, 1) {
		err := c.codec.Close()
		close(c.closeChan)
		if c.mgr != nil {
			if c.mgr.svr.onDisconnect != nil {
				c.mgr.svr.onDisconnect(c)
			}
			c.mgr.DeleteSession(c)
		}
		return err
	}

	return ErrClosedLock
}

// Send TODO
func (c *Session) Send(p interface{}, seq uint32) error {
	if c.isClosed() {
		return ErrSessionClosed
	}
	if c.sendChan == nil {
		return c.codec.Send(p, seq)
	}
	pack := &Packet{
		Seq: seq,
		Msg: p,
	}

	c.sendChan <- pack
	return nil
}

func (c *Session) sender() {
	defer c.Close()
	for {
		select {
		case pack := <-c.sendChan:
			if err := c.codec.Send(pack.Msg, pack.Seq); err != nil {
				xlog.Info(err)
				return
			}
		case <-c.closeChan:
			return
		}
	}
}

func (c *Session) isClosed() bool {
	return atomic.LoadInt32(&c.closed) == 1
}
