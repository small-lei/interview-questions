# 全排列。
```text
入参是一个数组，里面的数字不重复。返回的结果是一个数组的数组。入参 [1 2 3] 返回入参的全排列[[1 2 3][1 3 2][2 1 3][2 3 1][3 1 2][3 2 1]]
```
在 Go 语言中，可以使用回溯算法（Backtracking）来实现全排列。对于输入的数组 [1, 2, 3]，回溯算法通过递归地交换数组中的元素，生成所有可能的排列组合。
```go
package main

import "fmt"

// 生成全排列
func permute(nums []int) [][]int {
    var result [][]int
    backtrack(nums, 0, &result)
    return result
}

// 回溯函数
func backtrack(nums []int, start int, result *[][]int) {
    if start == len(nums) {
        // 深拷贝当前排列
        temp := make([]int, len(nums))
        copy(temp, nums)
        *result = append(*result, temp)
        return
    }
    
    for i := start; i < len(nums); i++ {
        // 交换元素
        nums[start], nums[i] = nums[i], nums[start]
        // 递归生成子排列
        backtrack(nums, start+1, result)
        // 回溯
        nums[start], nums[i] = nums[i], nums[start]
    }
}

func main() {
    nums := []int{1, 2, 3}
    result := permute(nums)
    fmt.Println(result)
}

```
#### 解释：
permute 函数：

接收一个数组作为参数，初始化结果数组 result，并调用 backtrack 函数来生成全排列。
backtrack 函数：

通过递归和交换数组元素来生成排列。
当递归到 start == len(nums) 时，表示找到了一个完整的排列，使用 make 和 copy 来保存这个排列的副本，然后将其加入结果数组中。
在递归过程中，每次将当前索引 start 位置的元素与后面的元素交换，然后递归处理剩余部分，递归结束后再交换回来（即“回溯”操作），恢复原来的数组顺序。
main 函数：

调用 permute 函数，传入 [1, 2, 3] 作为参数，并打印结果。
```text
[[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
```