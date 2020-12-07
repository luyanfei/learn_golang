package handle_test

import (
	"bytes"
	"errors"
	. "handle"
	"log"
	"strings"
	"testing"
)

func TestHandle(t *testing.T) {
	var (
		conf   = Config{Address: "China, ZheJiang"}
		buf    bytes.Buffer
		logger = log.New(&buf, "", 0)
		out    bytes.Buffer
	)
	WriteConfig(&out, &conf, logger)
	if !strings.Contains(out.String(), "China, ZheJiang") {
		t.Errorf("output does not contains address, the out is: %v.", out.String())
	}
	if len(buf.String()) != 0 {
		t.Errorf("log output is not empty: %v.", buf.String())
	}
}

type errWriter struct{}

func (e *errWriter) Write(p []byte) (int, error) {
	return 0, errors.New("error writer")
}

func TestHandleError(t *testing.T) {
	var (
		conf   = Config{Address: "China, ZheJiang"}
		buf    bytes.Buffer
		logger = log.New(&buf, "", 0)
		w      = errWriter{}
	)
	err := WriteConfig(&w, &conf, logger)
	if err == nil {
		t.Errorf("no error")
	}
	if len(buf.String()) == 0 {
		t.Errorf("no logger output")
	}
	if strings.Count(buf.String(), "error writer") != 1 {
		t.Errorf("only one 'error writer' is permitted, log message is: %s.", buf.String())
	}
}
