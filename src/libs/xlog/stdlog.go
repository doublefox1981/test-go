package xlog

import (
	"flag"
	"libs/xlog/lumberjack"
	"log"
	"os"
	"path/filepath"
)

var (
	std          *log.Logger
	stdWriter    *lumberjack.Logger
	logToConsole = true
	logLevel     = offLog
)

const (
	offLog int = iota
	debugLog
	infoLog
	warnLog
	errorLog
	fatalLog
)

var logName = []string{
	debugLog: "DEBUG",
	infoLog:  "INFO",
	warnLog:  "WARNING",
	errorLog: "ERROR",
	fatalLog: "FATAL",
}

func init() {
	var (
		fname string
		dir   string
	)

	exeName := filepath.Base(os.Args[0])

	flag.StringVar(&dir, "stdlogdir", "slogs", "std log path")
	flag.StringVar(&fname, "logname", exeName+".log", "log file name")
	flag.BoolVar(&logToConsole, "logtoconsole", true, "log to console")
	flag.IntVar(&logLevel, "loglevel", 0, "log level")
	fname = dir + "/" + fname

	if err := os.MkdirAll(dir, 0744); err != nil {
		panic("make dir logs")
	}

	stdWriter = &lumberjack.Logger{
		Filename:   fname,
		MaxSize:    100,
		MaxBackups: 1000,
		MaxAge:     1,
	}

	std = log.New(stdWriter, "", log.LstdFlags)
}

// Debug TODO
func Debug(v ...interface{}) {
	if logLevel > debugLog {
		return
	}
	std.Print(logName[debugLog], v)
	if logToConsole {
		log.Print(logName[debugLog], v)
	}
}

// Info TODO
func Info(v ...interface{}) {
	if logLevel > infoLog {
		return
	}
	std.Print(logName[infoLog], v)
	if logToConsole {
		log.Print(logName[infoLog], v)
	}
}

// Warn TODO
func Warn(v ...interface{}) {
	if logLevel > warnLog {
		return
	}
	std.Print(logName[warnLog], v)
	if logToConsole {
		log.Print(logName[warnLog], v)
	}
}

// Error TODO
func Error(v ...interface{}) {
	if logLevel > errorLog {
		return
	}
	std.Print(logName[errorLog], v)
	if logToConsole {
		log.Print(logName[errorLog], v)
	}
}

// Fatal TODO
func Fatal(v ...interface{}) {
	if logLevel > fatalLog {
		return
	}
	std.Print(logName[fatalLog], v)
	if logToConsole {
		log.Print(logName[fatalLog], v)
	}
}

// Sync TODO
func Sync() error {
	return stdWriter.Sync()
}
