# 介绍下gmp，go的gc 
```text
GMP 模型 (Goroutine-M-Processor)
GMP 是 Go 语言并发模型的核心架构，由三个关键组件组成：

G (Goroutine):

轻量级线程，Go 并发的基本单位
初始栈大小仅 2KB，远小于线程(通常MB级)
通过 go 关键字创建
M (Machine):

操作系统线程的抽象
由操作系统调度，真正执行代码的实体
一个 M 同一时间只能运行一个 G
P (Processor):

逻辑处理器，包含运行 Goroutine 的上下文
维护本地运行队列(runq)
默认数量等于 CPU 核心数(可通过 GOMAXPROCS 调整)

Go 的垃圾回收(GC)
Go 使用三色标记清除算法的并发垃圾回收器，主要特点：

分代假设优化：

虽然不严格分代，但通过写屏障(write barrier)实现部分代际优化
工作流程：

标记阶段：遍历可达对象(从栈/全局变量出发)
标记终止：STW(Stop The World)短暂暂停
清除阶段：回收不可达对象内存
关键参数：

GOGC：控制GC触发阈值(默认100%，即内存翻倍时触发)
GODEBUG=gctrace=1：输出GC详细日志
```

go语言中的内存逃逸是什么，会有什么影响，如何避免 
```text
内存逃逸是指本应在栈上分配的对象，由于生命周期超出当前函数作用域，被迫分配到堆上的现象。Go编译器通过逃逸分析(escape analysis)来确定变量的存储位置。
```
```go
func escape() *int {
    v := 42  // 逃逸到堆
    return &v
}

func closure() func() int {
    x := 10  // 逃逸到堆
    return func() int {
        return x
    }
}

ch := make(chan *int)
go func() {
    x := 100  // 逃逸到堆
    ch <- &x
}()
```
## 内存逃逸的影响
```text
1.性能影响：
堆分配比栈分配慢10-100倍
增加GC压力
降低缓存局部性

2.诊断方法：
go build -gcflags="-m -l" main.go
```