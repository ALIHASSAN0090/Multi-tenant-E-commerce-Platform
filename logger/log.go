package logger

import (
	"context"
	"ecommerce-platform/logger/log_repo"
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
)

type AppLogger struct {
	logger *log.Logger
	ctx    context.Context
}

type aboveCallerKey struct{}

func (l *AppLogger) WithAboveCaller(i int) log_repo.IAppLogger {
	return &AppLogger{
		logger: l.logger,
		ctx:    context.WithValue(l.ctx, aboveCallerKey{}, i),
	}
}

func (l *AppLogger) With(ctx context.Context) log_repo.IAppLogger {
	return &AppLogger{
		logger: l.logger,
		ctx:    ctx,
	}
}

func (l *AppLogger) WithValue(key, value interface{}) log_repo.IAppLogger {
	return &AppLogger{
		ctx:    context.WithValue(l.ctx, key, value),
		logger: l.logger,
	}
}

func (l *AppLogger) Info(args ...interface{}) {
	l.logger.Infoln(args...)
}

func (l *AppLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *AppLogger) Error(args ...interface{}) {
	l.logger.Errorln(args...)
}

func (l *AppLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *AppLogger) Warning(args ...interface{}) {
	l.logger.Warnln(args...)
}

func (l *AppLogger) Warningf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *AppLogger) Fatal(args ...interface{}) {
	l.logger.Fatalln(args...)
}

func (l *AppLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *AppLogger) Debug(args ...interface{}) {
	l.logger.Debugln(args...)
}

func (l *AppLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *AppLogger) Trace(args ...interface{}) {
	l.logger.Traceln(args...)
}

func (l *AppLogger) Report(args ...interface{}) {
	l.captureException(args...)
	l.logger.Errorln(args...)
}

func (l *AppLogger) Reportf(format string, args ...interface{}) {
	l.captureMsg(fmt.Sprintf(format, args...))
	l.logger.Errorf(format, args...)
}

func (l *AppLogger) captureException(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	msg = msg[:len(msg)-1]
	sentry.CaptureMessage(msg)
}

func (l *AppLogger) captureMsg(msg string) {
	sentry.CaptureMessage(msg)
}

func New() log_repo.IAppLogger {
	return NewWithDebugLevel(log.DebugLevel)
}

func NewWithDebugLevel(debugLevel log.Level) log_repo.IAppLogger {
	logger := log.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(debugLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
	logger.SetReportCaller(false)

	return &AppLogger{
		logger: logger,
		ctx:    context.Background(),
	}
}
