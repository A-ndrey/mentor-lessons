package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	//    ex1()
	//    ex2()
	//    ex3()
	ex4()
}

func ex1() {
	devLog := zap.Must(zap.NewDevelopment())

	devLog.Debug("debug msg", zap.Bool("bool val", true), zap.Time("time val", time.Now()))
	devLog.Info("info msg", zap.Bool("bool val", true), zap.Time("time val", time.Now()))
}

func ex2() {
	prodLog := zap.Must(zap.NewProduction())

	prodLog.Debug("debug msg", zap.Bool("bool val", true), zap.Time("time val", time.Now()))
	prodLog.Info("info msg", zap.Bool("bool val", true), zap.Time("time val", time.Now()))
}

func ex3() {
	logger := zap.Must(zap.NewProduction())
	sugar := logger.Sugar()

	sugar.Info("info msg", "bool val", true, "time val", time.Now())
	sugar.Infow("info msg", "bool val", true, "time val", time.Now())
	sugar.Infof("info msg bool val=%v, time val=%v", true, time.Now())
}

type st struct{}

func (st) Enabled(lvl zapcore.Level) bool {
	return lvl > zapcore.DebugLevel
}

func ex4() {
	logger := zap.Must(zap.NewProduction(zap.WithCaller(false), zap.AddStacktrace(st{}), zap.Fields(zap.String("function", "ex4"))))
	logger.Info("info msg")
	logger.Warn("warn msg")
}
