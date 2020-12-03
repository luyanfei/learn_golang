/*
Opaque errors: assert errros for behaviour, not type.
NetError接口的Temporary()方法用于判断error是否为临时错误。
请实现IsTemporary方法，当err实现了NetError接口，并且Temporary()方法返回为true的情况下，才会将err判定为临时错误。
*/
package temporary

type NetError interface {
	error
	Temporary() bool
}

func IsTemporary(err error) bool {
	te, ok := err.(NetError)
	return ok && te.Temporary()
}
