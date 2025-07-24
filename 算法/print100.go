package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i += 2 {
			<-ch1
			fmt.Println("协程1:", i)
			if i < 100 { // 确保协程2能收到最后一个信号
				ch2 <- struct{}{}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-ch2
			fmt.Println("协程2:", i)
			if i < 100 { // 确保协程1能收到最后一个信号
				ch1 <- struct{}{}
			}
		}
	}()

	// 启动第一个协程
	ch1 <- struct{}{}

	wg.Wait()
}
