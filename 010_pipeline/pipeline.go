package pipeline

func Generator(done <-chan interface{}, integers ...int) <-chan int {
	//将integers slice转换成channel，并且能够用close(done)来关闭该返回的channel。

}

func Multiply(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
	//将输入的intStream中的每个数乘上multiplier，返回新的channel，并且能够用close(done)来关闭该返回的channel。

}

func Add(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
	//将输入的intStream中的每个数加上additive，返回新的channel，并且能够用close(done)来关闭该返回的channel。

}
