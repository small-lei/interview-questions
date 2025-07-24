# go语言中有哪些情况会导致内存泄露
```go
1. 未关闭的资源
func leakyFile() {
    f, _ := os.Open("data.txt") // 泄漏点：未调用f.Close()
    // 使用文件...
}
数据库连接未释放
func leakyDB() {
    db, _ := sql.Open("mysql", "user:pwd@/dbname")
    rows, _ := db.Query("SELECT * FROM users") // 泄漏点：未调用rows.Close()
    // 处理结果...
}

2. 全局变量累积
缓存无限增长
var cache = make(map[string]interface{})

func addToCache(key string, value interface{}) {
cache[key] = value // 无清理机制会导致内存泄漏
}

3. Goroutine泄漏
无限阻塞的Goroutine
func leakyGoroutine() {
    ch := make(chan int)
    
    go func() {
        val := <-ch // 阻塞等待，永远不会被触发
        fmt.Println(val)
    }()

// 忘记关闭或发送数据到ch
}

未处理的channel
func worker(ch chan struct{}) {
    for {
        select {
            case <-ch:
            return
            default:
            // 持续工作...
            }
        }
    }
    
    func main() {
        ch := make(chan struct{})
        go worker(ch) // 泄漏点：没有终止机制
    }

4. 定时器未停止
time.Ticker未停止
func leakyTicker() {
    ticker := time.NewTicker(time.Second)
    for range ticker.C {
        // 处理逻辑...
        // 如果没有break条件且不调用ticker.Stop()会导致泄漏
    }
}

5. 循环引用
对象互相引用
type Node struct {
    next *Node
    data []byte
}

func circularRef() {
    n1 := &Node{data: make([]byte, 1024)}
    n2 := &Node{data: make([]byte, 1024)}
    n1.next = n2
    n2.next = n1 // 循环引用，即使外部不再引用也无法回收
}

检测与预防方法
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/heap

```