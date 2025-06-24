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
		var msg string
		var args []any
		if len(msgAndValues) > 0 {
			if v, ok := msgAndValues[0].(string); ok {
				msg = v
				args = msgAndValues[1:]
			} else {
				msg = "assertion failed"
			}
		}
		self.Log.Error(msg, args...)
		self.T.Fail()
	}
}

func (self Zester) Must(cond bool, msgAndValues ...any) {
	self.T.Helper()
	if !cond {
		var msg string
		var args []any
		if len(msgAndValues) > 0 {
			if v, ok := msgAndValues[0].(string); ok {
				msg = v
				args = msgAndValues[1:]
			} else {
				msg = "must condition failed"
			}
		}
		self.Log.Error(msg, args...)
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
