package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("排序前:", arr)
	sortedArr := quickSort(arr)
	fmt.Println("排序后:", sortedArr)
}
