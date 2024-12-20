package dao_service_impl

import (
	"context"
)

type IAppLogger interface {
	Trace(args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Report(args ...interface{})
	Reportf(format string, args ...interface{})
	With(ctx context.Context) IAppLogger
	WithValue(key, value interface{}) IAppLogger
	WithAboveCaller(int) IAppLogger
}
