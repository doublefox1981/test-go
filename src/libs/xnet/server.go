package xnet

import (
	"fmt"
	"libs/xlog"
	"net"
)

// Server TODO
type Server struct {
	listener     net.Listener
	protocol     Protocol
	mgr          *SessionMgr
	sndChanSize  int
	onConnect    func(c *Session) bool
	onDisconnect func(c *Session)
	onMsg        func(c *Session, p *Packet)
}

// NewTCPServer TODO
func NewTCPServer(addr string, proto Protocol, sndChanSize int) *Server {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		xlog.Error(fmt.Sprintf("listen error on %s, because %s", addr, err))
		return nil
	}
	xlog.Info(fmt.Sprintf("listen on tcp@%s", addr))
	s := &Server{
		listener:    l,
		protocol:    proto,
		sndChanSize: sndChanSize,
	}
	s.mgr = NewSessionMgr(s)
	return s
}

// Serve TODO
func (s *Server) Serve() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			xlog.Error(fmt.Sprintf("accept error on : %s, error : %v", s.listener.Addr().String(), err))
			return
		}
		codec := s.protocol.NewCodec(conn)
		if err != nil {
			conn.Close()
			continue
		}
		c := s.mgr.CreateSession(conn, codec, s.sndChanSize)
		var b = true
		if s.onConnect != nil {
			b = s.onConnect(c)
			if !b {
				c.Close()
				continue
			}
		}
		go s.ServeOne(c)
	}
}

// ServeOne TODO
func (s *Server) ServeOne(c *Session) {
	for {
		p, err := c.Receive()
		if err != nil {
			xlog.Info(err)
			c.Close()
			return
		}
		if s.onMsg != nil {
			s.onMsg(c, p)
		}
		s.protocol.PutPacket(p)
	}
}

// SetOnConnect TODO
func (s *Server) SetOnConnect(f func(c *Session) bool) {
	s.onConnect = f
}

// SetOnDisconnect TODO
func (s *Server) SetOnDisconnect(f func(c *Session)) {
	s.onDisconnect = f
}

// SetOnMsg TODO
func (s *Server) SetOnMsg(f func(c *Session, p *Packet)) {
	s.onMsg = f
}
