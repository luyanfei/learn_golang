package handling_test

import (
	"bytes"
	"errors"
	. "handling"
	"strings"
	"testing"
)

var st Status = Status{Code: 200, Reason: "OK"}
var headers []Header = []Header{
	{Key: "Accepted", Value: "text/html"},
	{Key: "Content-Type", Value: "text/html"},
}

func TestWriteResponse(t *testing.T) {
	var out bytes.Buffer
	var body string = "body content"
	bodyReader := strings.NewReader(body)
	err := WriteResponse(&out, st, headers, bodyReader)
	if err != nil {
		t.Errorf("there should be no error")
	}
	var expected string = "HTTP/1.1 200 OK\r\nAccepted: text/html\r\nContent-Type: text/html\r\n\r\nbody content"
	if out.String() != expected {
		t.Errorf("writed content is not expected")
	}
}

type errWriter struct{}

func (w *errWriter) Write(p []byte) (int, error) {
	return 0, errors.New("error writer")
}

func TestWriteResponseWithWriteError(t *testing.T) {
	w := errWriter{}
	var body string = "body content"
	bodyReader := strings.NewReader(body)
	err := WriteResponse(&w, st, headers, bodyReader)
	if err == nil {
		t.Errorf("there should be error")
	}
	if err.Error() != "error writer" {
		t.Errorf("expected 'error writer', but got: %s", err.Error())
	}
}

type errReader struct{}

func (r *errReader) Read(p []byte) (int, error) {
	return 0, errors.New("error reader")
}

func TestWriteResponseWithReadError(t *testing.T) {
	var out bytes.Buffer
	r := errReader{}
	err := WriteResponse(&out, st, headers, &r)
	if err == nil {
		t.Errorf("there should be error")
	}
	if err.Error() != "error reader" {
		t.Errorf("expected 'error reader', but got: %s", err.Error())
	}
}
