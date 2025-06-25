package zest

import (
	"fmt"
	"os"
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

	strpairs := make([]string, 1, (len(pairs)/2)+1)
	strpairs[0] = self.pairs
	for i := 0; i < len(pairs); i += 2 {
		value := fmt.Sprintf("%v", pairs[i+1])
		// Escape all control characters
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

const nl = "\n"

func (self Logger) log(level string, msg string, values ...any) {
	self.T.Helper()
	fmt.Fprint(os.Stderr, level, " ")
	fmt.Fprintf(os.Stderr, msg, values...)
	if self.pairs != "" {
		fmt.Fprint(os.Stderr, " ", self.pairs)
	}
	fmt.Fprintln(os.Stderr)
}

func (self Logger) Info(msg string, values ...any) {
	const inf = "INF"
	self.T.Helper()
	self.log(inf, msg, values...)
}

func (self Logger) Warn(msg string, values ...any) {
	const wrn = "WRN"
	self.T.Helper()
	self.log(wrn, msg, values...)
}

func (self Logger) Error(msg string, values ...any) {
	const err = "ERR"
	self.T.Helper()
	self.log(err, msg, values...)
}
