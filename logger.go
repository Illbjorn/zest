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

func (self *Logger) With(pairs ...any) *Logger {
	self.T.Helper()
	if len(pairs)%2 != 0 || len(pairs) == 0 {
		return self
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
	newPairs := strings.Join(strpairs, " ")

	if self.pairs != "" {
		return &Logger{
			T:     self.T,
			pairs: self.pairs + " " + newPairs,
		}
	}
	return &Logger{
		T:     self.T,
		pairs: strings.Join(strpairs, " "),
	}
}

func (self Logger) Info(msg string, values ...any) {
	self.T.Helper()
	self.T.Logf("INF "+msg, values...)
}

func (self Logger) Warn(msg string, values ...any) {
	self.T.Helper()
	self.T.Logf("WRN "+msg, values...)
}

func (self Logger) Error(msg string, values ...any) {
	self.T.Helper()
	self.T.Logf("ERR "+msg, values...)
}
