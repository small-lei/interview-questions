# Go 语言中 slice 的扩容机制与并发安全性
```text
Go 的 slice 扩容遵循以下规则（以 append 操作为例）：

容量计算：

新容量 = 当前容量 × 2（当原容量 < 1024）
新容量 = 当前容量 × 1.25（当原容量 ≥ 1024）
内存分配：

每次扩容会分配新的底层数组
旧数据被复制到新数组
旧数组由 GC 回收

关键结论
扩容会导致内存重新分配，影响性能，建议合理预分配容量
并发环境下必须使用同步原语保护 slice 操作
多个 goroutine 读取同一个 slice 是安全的（前提是没有并发写操作）
```

# go 如何通过函数改变全局slice？
在 Go 中，全局变量的作用域是包级别的，因此可以通过函数来修改全局变量，包含全局的 slice。当你传递一个 slice 作为参数时，它本身是一个引用类型，因此可以直接通过函数来修改它的内容。

### 方法一：使用指针
```go
package main

import (
	"fmt"
)

var globalSlice = []int{1, 2, 3}

func modifySlice(slice *[]int) {
	// 修改 slice 的值
	(*slice)[0] = 99
	// 追加新的元素
	*slice = append(*slice, 4, 5, 6)
	fmt.Println("Inside modifySlice, after append:", *slice)
}

func main() {
	fmt.Println("Before modifySlice:", globalSlice)

	// 传入 slice 的指针
	modifySlice(&globalSlice)

	fmt.Println("After modifySlice:", globalSlice)
}
```

### 方法二：返回新的 slice
```go
package main

import (
    "fmt"
)

var globalSlice = []int{1, 2, 3}

func modifySlice(slice []int) []int {
    // 修改 slice 的值
    slice[0] = 99
    // 追加新的元素
    slice = append(slice, 4, 5, 6)
    return slice
}

func main() {
    fmt.Println("Before modifySlice:", globalSlice)
    
    // 函数返回新的 slice
    globalSlice = modifySlice(globalSlice)
    
    fmt.Println("After modifySlice:", globalSlice)
}
```

### 总结：
```text
slice 是引用类型，直接修改其元素不需要传递指针。
如果需要通过函数修改 slice 的长度或重新分配内容，建议使用指针或返回修改后的 slice。
```

#  不同的Slice在复制和传值时，是深拷贝还是浅拷贝？
```text
浅拷贝：默认情况下，slice 的复制或传递是浅拷贝，多个 slice 共享同一个底层数组。
深拷贝：如果需要独立的底层数组，需要手动分配并使用 copy 函数复制元素。
```