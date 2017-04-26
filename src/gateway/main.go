package main

import (
	"libs/xlog"
	"libs/xnet"
	"pb"

	"fmt"
)

func onConnect(s *xnet.Session) bool {
	xlog.Info(fmt.Sprintf("new session, id=%d, %s", s.ID(), s.RemoteAddr()))
	return true
}

func onDisconnect(s *xnet.Session) {
	xlog.Info(fmt.Sprintf("close session, id=%d, %s", s.ID(), s.RemoteAddr()))
}

func onMsg(s *xnet.Session, p *xnet.Packet) {
	if p.Cmd == 1 {
		r := &pb.Ask{}
		r.A = "a"
		r.B = 1
		s.Send(r, p.Seq)
		xlog.Info(p.Cmd, p.Len, p.Seq)
	}
}

func main() {
	defer func() {
		xlog.Sync()
		xlog.ZapSync()
	}()
	protocol := xnet.Protobuf(4096)
	protocol.RegisterPB(1, (*pb.Ask)(nil), func() interface{} {
		return &pb.Ask{}
	})
	s := xnet.NewTCPServer("0.0.0.0:12315", protocol, 32)
	if s == nil {
		return
	}
	s.SetOnConnect(onConnect)
	s.SetOnDisconnect(onDisconnect)
	s.SetOnMsg(onMsg)
	s.Serve()
}
