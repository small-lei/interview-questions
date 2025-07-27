# 1.	一个池子内有 N 个打火机，这些打火机有的可以正常使用，有的已损坏。打火机使用后有一定几率损坏。如何快速从池子里取出一个可用的打火机？
在这个场景中，我们可以设计一个简单的打火机池子（LighterPool）来管理多个打火机，并实现一个从池子中快速取出可用打火机的方法。以下是一个使用 Golang 实现的示例：

实现步骤
定义打火机结构体：表示打火机的状态（可用或已损坏）。
定义打火机池结构体：管理打火机的集合，支持快速获取可用打火机。
实现取出打火机的方法：从池子中随机选择一个可用打火机并返回。
代码实现
go
复制代码
```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Lighter 表示一个打火机
type Lighter struct {
	IsAvailable bool // 打火机是否可用
}

// LighterPool 表示打火机池
type LighterPool struct {
	lighters []*Lighter
	mu       sync.Mutex
}

// NewLighterPool 创建一个新的打火机池
func NewLighterPool(size int) *LighterPool {
	pool := &LighterPool{
		lighters: make([]*Lighter, size),
	}

	for i := 0; i < size; i++ {
		pool.lighters[i] = &Lighter{
			IsAvailable: true, // 初始时所有打火机都可用
		}
	}

	return pool
}

// UseLighter 使用打火机，随机决定是否损坏
func (lp *LighterPool) UseLighter() *Lighter {
	lp.mu.Lock()
	defer lp.mu.Unlock()

	var availableLighters []*Lighter

	// 收集可用的打火机
	for _, lighter := range lp.lighters {
		if lighter.IsAvailable {
			availableLighters = append(availableLighters, lighter)
		}
	}

	// 如果没有可用的打火机，返回 nil
	if len(availableLighters) == 0 {
		return nil
	}

	// 随机选择一个可用的打火机
	selected := availableLighters[rand.Intn(len(availableLighters))]

	// 使用后决定是否损坏
	if rand.Float32() < 0.3 { // 30% 概率损坏
		selected.IsAvailable = false
		fmt.Println("打火机损坏了！")
	}

	return selected
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子

	pool := NewLighterPool(10) // 创建一个包含10个打火机的池子

	// 模拟使用打火机
	for i := 0; i < 15; i++ {
		if lighter := pool.UseLighter(); lighter != nil {
			fmt.Println("成功使用打火机！")
		} else {
			fmt.Println("没有可用的打火机。")
		}
		time.Sleep(500 * time.Millisecond) // 模拟使用间隔
	}
}

```

代码说明
```text
Lighter 结构体：包含一个布尔值 IsAvailable，指示打火机是否可用。
LighterPool 结构体：包含打火机的切片和一个互斥锁，用于确保并发安全。
NewLighterPool 函数：初始化一个打火机池，所有打火机初始状态为可用。
UseLighter 方法：
加锁以确保并发安全。
收集所有可用的打火机。
随机选择一个可用的打火机，并以 30% 的概率将其设置为不可用。
main 函数：模拟使用打火机的过程，打印结果。
运行示例
当你运行这个程序时，它会模拟多次使用打火机，输出可用和已损坏的打火机的信息。

注意事项
这里的概率和打火机的初始状态可以根据实际需求进行调整。
你可以在多线程场景中使用这个池子，sync.Mutex 确保了并发安全。
```