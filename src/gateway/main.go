package main

import (
	"fmt"
	"libs/xlog"
	"libs/xnet"
	"libs/xprofile"
	"net/http"
	"os"
	"os/signal"
	"pb"

	_ "net/http/pprof"

	"flag"

	_ "github.com/mkevac/debugcharts"
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

func signalf() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	for {
		s := <-c
		fmt.Println("signal", s)
		break
	}
}

func main() {
	flag.Parse()
	defer func() {
		xlog.Sync()
		xlog.ZapSync()
	}()
	defer xprofile.Start().Stop()
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
	go func() {
		xlog.Info(http.ListenAndServe("0.0.0.0:6060", nil))
	}()
	go s.Serve()
	signalf()
}
