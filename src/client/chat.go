package main

import (
	"libs/xlog"
	"libs/xnet"
	"pb"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	defer func() {
		xlog.Sync()
		xlog.ZapSync()
	}()
	protobuf := xnet.Protobuf(4096)
	protobuf.RegisterPB(1, (*pb.Ask)(nil), func() interface{} {
		return &pb.Ask{}
	})
	for i := 0; i < 100; i++ {
		go func() {
			cli, _ := xnet.Connect("tcp", "127.0.0.1:12315", protobuf, 32)
			a := &pb.Ask{
				A: "add",
				B: 1,
			}
			go func(c *xnet.Client) {
				for {
					p, err := cli.Receive()
					if err != nil {
						break
					}
					xlog.Info(p.Msg, err)
				}
			}(cli)
			for {
				err := cli.Send(a)
				if err != nil {
					xlog.Info(err)
					break
				}
				time.Sleep(time.Microsecond)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
