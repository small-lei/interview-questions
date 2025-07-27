

````text
Go 的 GMP 模型是什么？
G：Goroutine；

M：Machine，操作系统线程；

P：Processor，调度器调度队列（维护可执行G）；

一个P绑定一个M，运行G；

Go运行时负责调度G->M->P；

支持 work stealing、G 的复用等优化。
````

# GMP模型？
GMP 模型是 Go 语言中的一个核心并发模型，解释了 Goroutine（G）、操作系统线程（M）和处理器（P）之间的关系和调度机制。GMP 模型是 Go 语言并发编程的基础，它有助于理解 Goroutine 的调度、并发执行以及性能优化。以下是 GMP 模型的详细解释：

GMP 模型概述
Goroutine (G):
```text
定义: Goroutine 是 Go 语言的轻量级线程。每个 Goroutine 运行在用户级别，它们的创建和销毁开销较小。
特性: Goroutine 使用栈空间（通常非常小），可以在需要时动态增长或缩小。Goroutine 通过 Go 关键字 go 启动，例如：go func() { ... }()。
```

操作系统线程 (M):
```text
定义: 操作系统线程是由操作系统内核管理的线程。Go 语言运行时会将 Goroutine 映射到操作系统线程上。
特性: 操作系统线程的创建和销毁开销较大，相比 Goroutine 资源消耗更多。Go 运行时管理这些线程，控制 Goroutine 的调度。
```

处理器 (P):
```text
定义: 处理器是 Go 运行时的一个抽象，用于管理 Goroutine 的调度。一个 P 代表一个逻辑处理器，可以同时执行多个 Goroutine。
特性: 处理器负责为 Goroutine 提供运行的上下文和调度服务。Go 运行时通过 GOMAXPROCS 变量控制 P 的数量，从而限制同时执行的逻辑处理器数量。
```

GMP 模型的工作机制
调度:
```text
Goroutine 调度: Goroutine 由 Go 运行时调度，Goroutine 通过调度队列和调度器在不同的 M 和 P 上运行。
M 和 P 的映射: 多个 Goroutine 可以被映射到一个操作系统线程（M）上，同时一个 M 可能运行多个 Goroutine。当一个 Goroutine 被阻塞（例如等待 I/O 操作），调度器会将其他 Goroutine 分配到同一个 M 上。
```

GOMAXPROCS:
```text
定义: GOMAXPROCS 是一个环境变量或运行时设置，用于控制同时运行的 P 的数量。
作用: 通过限制 P 的数量，GOMAXPROCS 控制 Go 运行时的并发能力，从而影响程序的并发性能。
```

调度策略:
```text
G 运行时的调度器: Go 的调度器会根据 Goroutine 的状态（就绪、阻塞、等待）进行调度，以实现 Goroutine 的高效运行。
工作窃取算法: Go 运行时使用工作窃取算法来平衡负载。空闲的 P 可以从繁忙的 P 处窃取任务，以提高 CPU 使用率。
```

示例
以下是一个简单的示例，演示如何使用 Goroutine 来进行并发计算，并观察 GMP 模型的行为：
```go
go
复制代码
package main

import (
"fmt"
"runtime"
"time"
)

func main() {
runtime.GOMAXPROCS(2) // 设置同时运行的逻辑处理器数量为 2

    start := time.Now()

    for i := 0; i < 10; i++ {
        go func(i int) {
            time.Sleep(time.Second)
            fmt.Printf("Goroutine %d done\n", i)
        }(i)
    }

    // 等待所有 Goroutine 完成
    time.Sleep(5 * time.Second)

    fmt.Printf("Total time: %v\n", time.Since(start))
}
```

### 总结
```text
Goroutine (G): 轻量级线程，Go 语言的并发单元。
操作系统线程 (M): 执行 Goroutine 的实际线程，由操作系统管理。
处理器 (P): 逻辑处理器，管理 Goroutine 的调度。数量由 GOMAXPROCS 控制。
GMP 模型使得 Go 语言的并发编程高效且易于使用。通过理解 GMP 模型，你可以更好地优化 Go 程序的性能，并合理配置并发参数。
```