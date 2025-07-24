package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// 每次遍历将最大的元素冒泡到最后
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 交换相邻元素
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("排序前:", arr)
	sortedArr := bubbleSort(arr)
	fmt.Println("排序后:", sortedArr)
}
