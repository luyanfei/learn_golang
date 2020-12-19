package pipeline

func Generator(done <-chan interface{}, integers ...int) <-chan int {
	//将integers slice转换成channel，并且能够用close(done)来关闭该返回的channel。
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for _, v := range integers {
			select {
			case <-done:
				return
			case intStream <- v:
			}
		}
	}()
	return intStream
}

func Multiply(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
	//将输入的intStream中的每个数乘上multiplier，返回新的channel，并且能够用close(done)来关闭该返回的channel。
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for v := range intStream {
			select {
			case <-done:
				return
			case multipliedStream <- v * multiplier:
			}
		}
	}()
	return multipliedStream
}

func Add(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
	//将输入的intStream中的每个数加上additive，返回新的channel，并且能够用close(done)来关闭该返回的channel。
	additiveStream := make(chan int)
	go func() {
		defer close(additiveStream)
		for v := range intStream {
			select {
			case <-done:
				return
			case additiveStream <- v + additive:
			}
		}
	}()
	return additiveStream
}
