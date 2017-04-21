package main

import (
	"libs/xlog"
	"libs/xnet"
	"sync"
	"time"

	"go.uber.org/zap"
)

func main() {
	itemLogger := xlog.NewZapLogger("item.json", true)
	t1 := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1; j++ {
				itemLogger.Info("item", zap.String("field1", "11111"), zap.Int32("field2", 2))
				xlog.ZapInfo("exp", zap.String("field1", "11111"), zap.Int32("field2", 2))
			}
			wg.Done()
		}()
	}
	wg.Wait()
	t2 := time.Now()
	xlog.Info(t2.Sub(t1))

	protocol := xnet.Protobuf(4096)
	s := xnet.NewTCPServer("0.0.0.0:12315", protocol)
	if s == nil {
		return
	}
	s.Serve()
}
