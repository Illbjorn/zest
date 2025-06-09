package zest

import (
	"fmt"
	"strings"
	"testing"
)

type Logger struct {
	T     *testing.T
	pairs string
}

func (self *Logger) With(pairs ...any) {
	self.T.Helper()
	if len(pairs)%2 != 0 || len(pairs) == 0 {
		return
	}

	strpairs := make([]string, 0, len(pairs)/2)
	for i := 0; i < len(pairs); i += 2 {
		value := fmt.Sprintf("%v", pairs[i+1])
		// Replace all control characters with literals
		value = strings.ReplaceAll(value, "\n", "\\n")
		value = strings.ReplaceAll(value, "\t", "\\t")
		value = strings.ReplaceAll(value, "\v", "\\v")
		value = strings.ReplaceAll(value, "\f", "\\f")
		value = strings.ReplaceAll(value, "\r", "\\r")
		strpairs = append(strpairs, fmt.Sprintf("[%s]=>[%s]", pairs[i], value))
	}

	*self = Logger{
		T:     self.T,
		pairs: strings.Join(strpairs, " "),
	}
}

func (self Logger) Info(msg string, values ...any) {
	self.T.Helper()
	msg = fmt.Sprintf(msg, values...)
	if self.pairs != "" {
		msg += self.pairs
	}
	self.T.Log(msg)
}

func (self Logger) Error(msg string, values ...any) {
	self.T.Helper()
	msg = fmt.Sprintf(msg, values...)
	if self.pairs != "" {
		msg += self.pairs
	}
	self.T.Error(msg)
}
