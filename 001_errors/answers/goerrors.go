/*
1. 创建error interface
2. 创建errorString struct
3. 绑定到errorString struct上的Error()方法
4. 返回新errorString对象的New()方法
*/
package goerrors

type error interface {
	Error() string
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}
