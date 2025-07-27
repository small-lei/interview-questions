# 双数组求第k小
```text
题目描述
给定2个有序数组，均从小到打排序，求在两个数组中第k小的元素的值

例如：
a = [1, 3, 5, 7, 9]

b = [2, 4, 6, 8, 10]

返回第5小的数字， 为6
```

### 代码实现
```go
package main

import (
	"fmt"
)

// findKthSmallest 在两个有序数组中找到第 k 小的元素
func findKthSmallest(a []int, b []int, k int) int {
	m, n := len(a), len(b)

	// 边界条件，如果一个数组为空，直接返回另一个数组中第k小的元素
	if m == 0 {
		return b[k-1]
	}
	if n == 0 {
		return a[k-1]
	}

	// 如果 k == 1，返回两个数组中最小的第一个元素
	if k == 1 {
		if a[0] < b[0] {
			return a[0]
		}
		return b[0]
	}

	// 取k/2对应的下标 (如果数组长度不足k/2，取数组末尾)
	i := min(m, k/2) - 1
	j := min(n, k/2) - 1

	// 比较 a[i] 和 b[j]
	if a[i] <= b[j] {
		// a[i] 之前的元素（包括a[i]）不可能是第k小的元素，排除
		return findKthSmallest(a[i+1:], b, k-i-1)
	} else {
		// b[j] 之前的元素（包括b[j]）不可能是第k小的元素，排除
		return findKthSmallest(a, b[j+1:], k-j-1)
	}
}

// min 返回两个整数中的最小值
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	a := []int{1, 3, 5, 7, 9}
	b := []int{2, 4, 6, 8, 10}

	k := 5
	result := findKthSmallest(a, b, k)
	fmt.Printf("第%d小的数字是: %d\n", k, result)
}

```
