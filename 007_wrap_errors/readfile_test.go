package readfile_test

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	. "readfile"
	"strings"
	"testing"
)

func TestReadConfig(t *testing.T) {
	_, err := ReadConfig("./none.txt")
	var buf bytes.Buffer
	if err == nil {
		t.Errorf("err should not be nil.")
	}
	fmt.Fprintf(&buf, "original error: %T %v\n", errors.Cause(err), errors.Cause(err))
	fmt.Fprintf(&buf, "stack trace:\n%+v\n", err)
	if !strings.Contains(buf.String(), "open failed") {
		t.Errorf("open failed is not wraped.")
	}
	if !strings.Contains(buf.String(), "could not read config") {
		t.Errorf("ReadConfig message is not contained.")
	}
}
