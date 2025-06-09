package zest

import (
	"testing"
)

func New(t *testing.T) Zester {
	return Zester{
		T:      t,
		Helper: t.Helper,
		Logger: Logger{
			T: t,
		},
	}
}

type Zester struct {
	T *testing.T
	Logger
	Helper func()
}

func (self Zester) Assert(cond bool, msg string, values ...any) {
	self.Helper()
	if !cond {
		self.Error(msg, values...)
	}
}
