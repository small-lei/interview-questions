package main

import "fmt"

/*
defer 的执行顺序？
多个 defer 会 先进后出（LIFO）；

defer 在 return 之前执行，但绑定变量时机是 defer声明时，不是 return 时。
*/

func main() {
	var i int

	defer func() {
		i++
		fmt.Println("func1", i)
	}()

	defer func() {
		i++
		fmt.Println("func2", i)
	}()

	defer func() {
		i++
		fmt.Println("func3", i)
	}()
}
