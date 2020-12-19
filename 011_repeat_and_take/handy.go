package handy

func Repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	//TODO:反复地生成values slice的值，放入返回的channel，并且能够用close(done)来关闭该返回的channel。
}

func Take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	//TODO:从valueStream中抽取前num个元素，组成新的channel，并且能够用close(done)来关闭返回的channel。
}
