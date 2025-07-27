# go 实现快排

在 Go 中实现快速排序（QuickSort）是一个经典的算法问题。快速排序的核心思想是分治法，通过选取一个基准点（pivot），然后将数组分成两个子数组，一个部分的元素都小于或等于基准点，另一个部分的元素都大于基准点，最后对两个子数组递归地进行排序。

实现步骤
选择一个基准元素（通常选择第一个、最后一个、中间一个或随机选一个）。
通过一趟排序将数组分割成两个部分，左边部分都小于等于基准元素，右边部分都大于基准元素。
递归地对左右两个部分分别进行快速排序。
递归结束后，数组有序。
```go
package main

import (
    "fmt"
)

// partition 函数负责将数组进行分区，返回基准点的位置
func partition(arr []int, low, high int) int {
    pivot := arr[high] // 选择最后一个元素作为基准点
    i := low - 1       // i 是小于等于 pivot 的最后一个元素的索引

    for j := low; j < high; j++ {
        if arr[j] <= pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i] // 交换元素
        }
    }
    arr[i+1], arr[high] = arr[high], arr[i+1] // 把基准点放到正确的位置
    return i + 1 // 返回基准点的索引
}

// quickSort 递归排序函数
func quickSort(arr []int, low, high int) {
    if low < high {
        // 获取基准点的索引
        pi := partition(arr, low, high)

        // 递归对基准点左侧进行排序
        quickSort(arr, low, pi-1)
        // 递归对基准点右侧进行排序
        quickSort(arr, pi+1, high)
    }
}

func main() {
    arr := []int{10, 80, 30, 90, 40, 50, 70}
    fmt.Println("Unsorted array:", arr)

    quickSort(arr, 0, len(arr)-1)

    fmt.Println("Sorted array:", arr)
}
```
```text
代码解析
partition 函数：负责将数组划分为两部分。选择最后一个元素作为基准点，然后遍历数组，将小于等于基准点的元素移动到数组左侧，最后将基准点放到它正确的位置，返回基准点的索引。
quickSort 函数：递归调用，负责将数组的两部分分别进行排序。
如果 low < high，说明数组还需要排序。
partition 函数将数组划分为左右两部分，然后分别对这两部分递归调用 quickSort。

快速排序的时间复杂度
最优时间复杂度：O(n log n)，当每次选择的基准点都能将数组平衡地划分为两部分时，效率最高。
平均时间复杂度：O(n log n)，通常情况下，快速排序的平均时间复杂度较好。
最差时间复杂度：O(n^2)，当每次选择的基准点都不平衡（比如总是选择最大的或最小的元素）时，排序效率最差。
快速排序的空间复杂度
空间复杂度：O(log n)（递归栈空间），最优和平均情况下由于递归调用深度为 log n，但在最差情况下，递归深度会达到 n，空间复杂度为 O(n)。
总结
快速排序是一种高效的排序算法，平均时间复杂度为 O(n log n)，实现简单，广泛用于各种实际场景。不过，快速排序的最差时间复杂度为 O(n^2)，需要通过选择合适的基准点来避免最差情况发生。
```