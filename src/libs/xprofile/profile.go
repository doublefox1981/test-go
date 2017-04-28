package xprofile

import (
	"flag"
	"fmt"
	"libs/xlog"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"strings"
	"sync/atomic"
)

// Profiler TODO
type Profiler struct {
	closer         []func()
	memProfileRate int
	hookShutdown   bool
	stopped        uint32
}

var mode string

func init() {
	flag.StringVar(&mode, "profile", "cpu", "profile mode, such as \"cpu,mem,mutex,block,trace\"")
}

func (p *Profiler) cpuProfile() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		xlog.Info(fmt.Sprintf("pfofile: cound not create cpu pfofile %v", err))
		return
	}
	xlog.Info("profile: cpu profiling enabled")
	pprof.StartCPUProfile(f)
	p.closer = append(p.closer, func() {
		pprof.StopCPUProfile()
		f.Close()
		xlog.Info("profile: cpu profiling disabled")
	})
}

func (p *Profiler) memProfile() {
	f, err := os.Create("mem.pprof")
	if err != nil {
		xlog.Info(fmt.Sprintf("pfofile: cound not create memory pfofile %v", err))
		return
	}
	xlog.Info("profile: memory profiling enabled")
	old := runtime.MemProfileRate
	runtime.MemProfileRate = p.memProfileRate
	p.closer = append(p.closer, func() {
		pprof.Lookup("heap").WriteTo(f, 0)
		f.Close()
		runtime.MemProfileRate = old
		xlog.Info("profile: memory profiling disabled")
	})
}

func (p *Profiler) mutexProfile() {

}

func (p *Profiler) blockProfile() {

}

func (p *Profiler) traceProfile() {

}

// MemProfileRate TODO
func MemProfileRate(rate int) func(*Profiler) {
	return func(p *Profiler) {
		p.memProfileRate = rate
	}
}

// ShutdownHook TODO
func ShutdownHook(hook bool) func(*Profiler) {
	return func(p *Profiler) {
		p.hookShutdown = hook
	}
}

// Start TODO
func Start(options ...func(*Profiler)) *Profiler {
	var prof Profiler
	for _, option := range options {
		option(&prof)
	}
	mmodes := map[string]func(){
		"cpu":   prof.cpuProfile,
		"mem":   prof.memProfile,
		"mutex": prof.mutexProfile,
		"block": prof.blockProfile,
		"trace": prof.traceProfile,
	}
	modes := strings.FieldsFunc(mode, func(r rune) bool {
		return uint32(r) == ','
	})
	for _, m := range modes {
		f, ok := mmodes[m]
		if ok {
			f()
		}
	}

	if prof.hookShutdown {
		go func() {
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			<-c
			prof.Stop()
			os.Exit(0)
		}()
	}
	return &prof
}

// Stop TODO
func (p *Profiler) Stop() {
	if !atomic.CompareAndSwapUint32(&p.stopped, 0, 1) {
		return
	}
	for _, c := range p.closer {
		c()
	}
}
