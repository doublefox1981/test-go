# 对标准库log的改造

lumberjack/lumberjack.go 来源至 https://github.com/natefinch/lumberjack/blob/v2.0/lumberjack.go，实现了对日志文件按照大小和日期进行日志轮转

stdlog使用的标准库中的log包，整合在stdlog.go，可进行一些简单、低频的非结构化日志的记录，默认单个文件最大为100M，可通过命令行参数修改文件名、大小限制等，使用案例如下。
```go
package main

import (
	"libs/xlog"
)

func main() {
	xlog.Debug("hello")
	xlog.Info("hello")
	xlog.Warn("hello")
	xlog.Error("hello")
	xlog.Fatal("hello")
}
````


# 基于 uber zap 日志系统的整合
整合在zaplog.go，定位：结构化日志（一般为业务日志，序列化成json）的记录，该库不使用反射，效率高，结合lumberjack可实现日志轮转。json格式便于第三方工具例如elk套件，logtail等对日志进行收集，集中处理，使用案例如下。
```go
package main

import (
	"libs/xlog"
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
			for j := 0; j < 100000; j++ {
				itemLogger.Info("item", zap.String("field1", "11111"), zap.Int32("field2", 2))
				xlog.ZapInfo("exp", zap.String("field1", "11111"), zap.Int32("field2", 2))
			}
			wg.Done()
		}()
	}
	wg.Wait()
	t2 := time.Now()
	xlog.Info(t2.Sub(t1))
}
```