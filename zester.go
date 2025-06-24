package zest

import (
	"testing"
)

func New(t *testing.T) Zester {
	return Zester{
		T: t,
		Log: &Logger{
			T: t,
		},
	}
}

type Zester struct {
	T   *testing.T
	Log *Logger
}

func (self Zester) Assert(cond bool, msgAndValues ...any) {
	self.T.Helper()
	if !cond {
		if msg, ok := msgAndValues[0].(string); ok {
			self.T.Errorf(msg, msgAndValues[1:]...)
			return
		} else {
			self.T.Error("assertion failed")
		}
	}
}

func (self Zester) Must(cond bool, msg string, values ...any) {
	self.T.Helper()
	if !cond {
		self.T.Errorf(msg, values...)
		self.T.FailNow()
	}
}

func (self Zester) NoError(err error, msgAndValues ...any) {
	self.T.Helper()
	if len(msgAndValues) > 0 {
		self.Assert(err == nil, msgAndValues...)
	} else {
		self.Assert(err == nil, "expected no error, got [%s]", err)
	}
}
