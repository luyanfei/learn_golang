package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
使用goroutine要遵守以下原则：
If a goroutine is responsible for creating a goroutine, it is also responsible for ensuring it can stop the goroutine.
下面的程序在执行时会有内存的泄漏，newRandStream内部的匿名goroutine无法正常退出，程序运行时不会输出"newRandStream closure exited."。
请修改程序，令其能够正确结束newRandStream内部的匿名goroutine，并输出"newRandStream closure exited."
本题没有测试用例，用go run randstream.go来观察结果。(Concurrency in Go P93)
*/
func main() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}
	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	time.Sleep(1 * time.Second)
}
