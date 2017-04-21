package xlog

import (
	"libs/xlog/lumberjack"
	"os"
	"path"
	"path/filepath"

	"flag"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logDir               = "zlogs"
	productionMod        = false
	consoleZapLogger     *zap.Logger
	defaultFileZapLogger *zap.Logger
	defaultFileName      string
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

func init() {
	flag.StringVar(&logDir, "zaplogdir", "zlogs", "zap logs path")
	exeName := filepath.Base(os.Args[0])
	defaultFileName = exeName + ".json"

	err := os.MkdirAll(logDir, 0744)
	if err != nil {
		panic("make jsonlogs dir")
	}

	newDefaultLogger()
}

func newRotateWriter(fileName string) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   path.Join(logDir, fileName),
		MaxSize:    100,
		MaxAge:     1,
		MaxBackups: 1000,
	})
}

func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
	}
}

func NewZapLogger(fileName string, encodeAsJSON bool) *zap.Logger {
	encCfg := newEncoderConfig()
	var encoder zapcore.Encoder
	if encodeAsJSON {
		encoder = zapcore.NewJSONEncoder(encCfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(encCfg)
	}

	output := newRotateWriter(fileName)

	return zap.New(zapcore.NewCore(encoder, output, zap.NewAtomicLevel()))
}

func newDefaultLogger() {
	defaultFileZapLogger = NewZapLogger(defaultFileName, true)
	encCfg := newEncoderConfig()
	encoder := zapcore.NewConsoleEncoder(encCfg)
	consoleZapLogger = zap.New(zapcore.NewCore(encoder, os.Stdout, zap.NewAtomicLevel()))
}

func ZapDebug(msg string, fields ...zapcore.Field) {
	consoleZapLogger.Debug(msg, fields...)
	defaultFileZapLogger.Debug(msg, fields...)
}

func ZapInfo(msg string, fields ...zapcore.Field) {
	consoleZapLogger.Info(msg, fields...)
	defaultFileZapLogger.Info(msg, fields...)
}

func ZapWarn(msg string, fields ...zapcore.Field) {
	consoleZapLogger.Warn(msg, fields...)
	defaultFileZapLogger.Warn(msg, fields...)
}

func ZapError(msg string, fields ...zapcore.Field) {
	consoleZapLogger.Error(msg, fields...)
	defaultFileZapLogger.Error(msg, fields...)
}
