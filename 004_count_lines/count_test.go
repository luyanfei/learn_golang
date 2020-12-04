package count_test

import (
	. "count"
	"errors"
	"strings"
	"testing"
)

var proverbs string = `
Errors are values.
Don’t just check errors, handle them gracefully.
Don’t panic.
Make the zero value useful.
The bigger the interface, the weaker the abstraction.
interface{} says nothing.
Gofmt’s style is no one’s favorite, yet gofmt is everyone’s favorite.
Documentation is for users.
A little copying is better than a little dependency.
Clear is better than clever.
Concurrency is not parallelism.
Don’t communicate by sharing memory, share memory by communicating.
Channels orchestrate; mutexes serialize.

`

type countReader int

func (r countReader) Read(p []byte) (int, error) {
	return 0, errors.New("No EOF")
}

func TestCountLines(t *testing.T) {
	r := strings.NewReader(proverbs)
	lines, err := CountLines(r)
	if lines != 15 {
		t.Errorf("CountLines is wrong, expected value is 16, but got: %d", lines)
	}
	if err != nil {
		t.Errorf("there should be no err.")
	}

}

func TestCountLinesWithError(t *testing.T) {
	var c countReader = 0
	_, err := CountLines(c)
	if err == nil {
		t.Errorf("there should be error")
	}
}
