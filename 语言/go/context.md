# context的超时控制如何实现的
```text
Go 的 context 包通过以下方式实现超时控制：

数据结构：

timerCtx 是核心结构，嵌入了 cancelCtx 并增加了定时器
包含截止时间(deadline)和 time.Timer
关键方法：

WithTimeout/WithDeadline 创建带超时的 context
内部启动 goroutine 监听定时器
```
```go
// context包中的核心实现
type timerCtx struct {
    cancelCtx
    timer *time.Timer // 定时器
    deadline time.Time // 截止时间
}

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
    return WithDeadline(parent, time.Now().Add(timeout))
}

func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
    // ... 父context检查等逻辑
    
    c := &timerCtx{
        cancelCtx: newCancelCtx(parent),
        deadline:  d,
    }
    
    // 计算超时时间差
    dur := time.Until(d)
    if dur <= 0 {
        c.cancel(true, DeadlineExceeded) // 已超时
        return c, func() { c.cancel(false, Canceled) }
    }
    
    // 启动定时器
    c.timer = time.AfterFunc(dur, func() {
        c.cancel(true, DeadlineExceeded)
    })
    
    // ... 父context监听等逻辑
}


func doWork(ctx context.Context) {
select {
case <-time.After(5 * time.Second):
fmt.Println("工作完成")
case <-ctx.Done():
fmt.Println("工作取消:", ctx.Err())
}
}

func main() {
// 设置1秒超时
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel()

go doWork(ctx)

// 等待工作完成或取消
time.Sleep(2 * time.Second)
}
```
## 实现特点
```text
高效通知：

通过关闭 Done() channel 广播取消信号
所有监听该 channel 的 goroutine 会立即收到通知
级联取消：

父 context 取消会触发所有子 context 取消
通过 context 树形结构实现
资源回收：

超时或手动取消都会停止定时器
防止 goroutine 和定时器泄漏
最佳实践
总是使用 defer cancel() 确保资源释放
将 context 作为函数第一个参数传递
对长时间运行的操作检查 ctx.Done()
这种设计使得 Go 程序的超时控制既高效又易于理解，成为并发编程的重要模式。
```