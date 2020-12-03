package temporary_test

import (
	"errors"
	. "temporary"
	"testing"
)

type temporaryError struct {
	s    string
	flag bool
}

func (err temporaryError) Error() string {
	return err.s
}

func (err temporaryError) Temporary() bool {
	return err.flag
}

func TestIsTemporary(t *testing.T) {
	err1 := errors.New("test1")
	if IsTemporary(err1) {
		t.Errorf("plain error is temporary")
	}

	err2 := temporaryError{s: "temporary message", flag: true}
	if !IsTemporary(err2) {
		t.Errorf("ture temporary error is not temporary")
	}

	err3 := temporaryError{s: "temporary message", flag: false}
	if IsTemporary(err3) {
		t.Errorf("false temporary error is temporary")
	}
}
