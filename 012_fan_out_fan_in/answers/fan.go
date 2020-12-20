package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

//repeatFn返回的channel中，会不断写入由fn生成的值，直到用close(done)关闭。
func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

//primeFinder中素数的判断非常耗时，用于模拟需要大量计算的任务。
//primeFinder生成的channel中会是intStream中的素数。
func primeFinder(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
	primeStream := make(chan interface{})
	isPrime := func(n int) bool {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
	go func() {
		defer close(primeStream)
		for v := range intStream {
			if isPrime(v) {
				select {
				case <-done:
					return
				case primeStream <- v:
				}
			}
		}
	}()
	return primeStream
}

//将chan interface{}转换成chan int.
func toInt(done <-chan interface{}, inputStream <-chan interface{}) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for v := range inputStream {
			if iv, ok := v.(int); ok {
				select {
				case <-done:
					return
				case intStream <- iv:
				}
			}
		}
	}()
	return intStream
}

//获取valueStream中的前num个元素。
func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})

	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()
	return multiplexedStream
}

/*
main函数中用管道从不断生成的随机数中提取前10个素数，由于素数算法耗时较长，程序运行需要较长时间。
请用Fan-Out,Fan-In模式改写main，充分利用多核，以提高其运行效率。
*/
func main() {
	rand := func() interface{} {
		return rand.Intn(200000000)
	}
	done := make(chan interface{})
	defer close(done)

	start := time.Now()
	randIntStream := toInt(done, repeatFn(done, rand))
	fmt.Println("Primes:")
	numFinders := runtime.NumCPU()
	finders := make([]<-chan interface{}, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}
	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}
