package goerrors_test

import (
	"goerrors"
	"testing"
)

func TestNewEqual(t *testing.T) {
	// Different allocations should not be equal.
	if goerrors.New("abc") == goerrors.New("abc") {
		t.Errorf(`New("abc") == New("abc")`)
	}
	if goerrors.New("abc") == goerrors.New("xyz") {
		t.Errorf(`New("abc") == New("xyz")`)
	}

	// Same allocation should be equal to itself (not crash).
	err := goerrors.New("jkl")
	if err != err {
		t.Errorf(`err != err`)
	}
}
