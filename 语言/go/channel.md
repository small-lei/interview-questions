# Channel的优雅关闭，如何避免关闭一个已经关闭的channel
## 1. 标准解决方案（类似context的cancel机制）
```go
// 使用sync.Once确保只关闭一次（类似context的cancel只执行一次）
type SafeChannel struct {
    ch    chan int
    once  sync.Once
    closed bool
}

func (sc *SafeChannel) Close() {
    sc.once.Do(func() {
        close(sc.ch)
        sc.closed = true
    })
}

// 使用示例（类似链表操作中的dummy节点保护）
func main() {
    sc := &SafeChannel{ch: make(chan int, 10)}
    // ... 生产消费逻辑 ...
    sc.Close() // 多次调用也只会关闭一次
    sc.Close() // 安全
}
```
## 2.通道包装器模式（类似context的树形结构管理
```go
// 通道管理器（类似context的树形取消传播）
type ChannelManager struct {
    mu      sync.Mutex
    chs     map[chan int]struct{}
    closed  bool
}

func (cm *ChannelManager) NewChannel() chan int {
    cm.mu.Lock()
    defer cm.mu.Unlock()
    
    if cm.closed {
        return nil
    }
    ch := make(chan int, 10)
    cm.chs[ch] = struct{}{}
    return ch
}

func (cm *ChannelManager) CloseAll() {
    cm.mu.Lock()
    defer cm.mu.Unlock()
    
    if !cm.closed {
        for ch := range cm.chs {
            close(ch)
        }
        cm.closed = true
    }
}
```