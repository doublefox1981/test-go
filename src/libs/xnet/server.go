package xnet

import (
	"fmt"
	"libs/xlog"
	"net"
)

// Server TODO
type Server struct {
	listener   net.Listener
	protocol   Protocol
	mgr        *SessionMgr
	onConnect  func(c net.Conn) bool
	msgHandler func(c *Session, p *Packet)
}

// NewTCPServer TODO
func NewTCPServer(addr string, proto Protocol) *Server {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		xlog.Error(fmt.Sprintf("listen error on %s, because %s", addr, err))
		return nil
	}
	xlog.Info(fmt.Sprintf("listen on tcp@%s", addr))
	return &Server{
		listener: l,
		protocol: proto,
		mgr:      NewSessionMgr(),
	}
}

// Serve TODO
func (s *Server) Serve() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			xlog.Error(fmt.Sprintf("accept error on : %s, error : %v", s.listener.Addr().String(), err))
			return
		}
		var b = true
		if s.onConnect != nil {
			b = s.onConnect(conn)
			if !b {
				conn.Close()
				continue
			}
		}
		codec := s.protocol.NewCodec(conn)
		if err != nil {
			conn.Close()
			continue
		}
		c := s.mgr.CreateSession(codec)
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
		if s.msgHandler != nil {
			s.msgHandler(c, p)
		}
		s.protocol.PutPacket(p)
	}
}

//
func (s *Server) SetOnConnect(f func(c net.Conn) bool) {
	s.onConnect = f
}
