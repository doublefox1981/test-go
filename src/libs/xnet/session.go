package xnet

import (
	"sync/atomic"
)

var allocSessionID uint64

// Session TODO
type Session struct {
	id    uint64
	codec Codec
	mgr   *SessionMgr
}

func newSession(mgr *SessionMgr, codec Codec) *Session {
	c := &Session{
		codec: codec,
		mgr:   mgr,
		id:    atomic.AddUint64(&allocSessionID, 1),
	}

	return c
}

// Receive TODO
func (c *Session) Receive() (*Packet, error) {
	return c.codec.Receive()
}

// Close TODO
func (c *Session) Close() error {
	c.codec.Close()
	c.mgr.DeleteSession(c)
	return nil
}

// Send TODO
func (c *Session) Send(p interface{}) error {
	return c.codec.Send(p)
}
