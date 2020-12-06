package handling

import (
	"fmt"
	"io"
)

type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

/*
下面的WriteResponse函数中关于写入的错误判断太啰嗦了，通过添加适当的结构及方法，可避免每次写入时都要判断错误的写法。请改写WriteResponse的实现。
func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
	if err != nil {
		return err
	}

	for _, h := range headers {
		_, err := fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.Value)
		if err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, "\r\n"); err != nil {
		return err
	}

	_, err = io.Copy(w, body)
	return err
}
*/

type errorWriter struct {
	io.Writer
	err error
}

func (e *errorWriter) Write(p []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	var n int
	n, e.err = e.Writer.Write(p)
	return n, nil
}

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &errorWriter{Writer: w}
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)

	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}

	fmt.Fprint(ew, "\r\n")
	//io.Reader error should be processed too
	if _, err := io.Copy(ew, body); err != nil {
		return err
	}
	return ew.err
}
