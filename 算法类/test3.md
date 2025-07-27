# 题目描述
输入m, n 输出一个m行n列的蛇形矩阵，m,n 的取值范围均为[2, 10], 2<=m, n <=10;
```text
示例：
输入 m=3, n =4
输出
1 2 6 7
3 5 8 11
4 9 10 12
```

### 解答
要生成一个蛇形矩阵，可以按如下规则填充矩阵：
```text
从左到右填充第一行。
从上到下填充第二列。
从右到左填充最后一行。
从下到上填充第一列。
继续按这个顺序迭代，直到矩阵被完全填满。
```

```go
package main

import (
	"fmt"
)

func generateSnakeMatrix(m, n int) [][]int {
	// 初始化 m 行 n 列的矩阵
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// 定义方向和边界
	top, bottom, left, right := 0, m-1, 0, n-1
	num := 1

	for top <= bottom && left <= right {
		// 从左到右填充
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		// 从上到下填充
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		// 从右到左填充
		if top <= bottom {
			for i := right; i >= left; i-- {
				matrix[bottom][i] = num
				num++
			}
			bottom--
		}

		// 从下到上填充
		if left <= right {
			for i := bottom; i >= top; i-- {
				matrix[i][left] = num
				num++
			}
			left++
		}
	}

	return matrix
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

func main() {
	m, n := 3, 4
	matrix := generateSnakeMatrix(m, n)
	printMatrix(matrix)
}

```

### 代码解析：
```text
generateSnakeMatrix：用于生成蛇形矩阵。通过变量 top, bottom, left, right 来控制当前的边界，并依次按照四个方向填充矩阵。
方向的填充顺序：
从左到右填充矩阵的上边界。
从上到下填充右边界。
从右到左填充下边界。
从下到上填充左边界。
```


