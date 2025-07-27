# 代码题，实现一个本地的cache ，达到线程安全?
```go
type MyMap interface {
	Add(key string, value int)
	Remove(key string)
	Get(key string) (int, bool)
	SnapShot()map[string]int
}
```
说，现在这个缓存的key很多，有几千万个，这个snapShot如果按你刚才的思路（深拷贝），就会阻塞时间很长，有什么办法优化？

### 解答
为了解决在处理几千万个键值对时 SnapShot() 进行深拷贝会导致阻塞时间长的问题，可以使用读写锁分片或Copy-On-Write 思路来优化。在这种场景下，快照的性能至关重要，因此你可以通过减少锁的争用、提高并发性和分片处理来优化。

实现思路

分片设计（Sharded Map）：

将整个 map 按键的哈希分片成多个独立的部分，每个分片都有自己的锁。这种方式可以显著减少锁的竞争，提高并发性能。
快照操作时，可以逐个分片进行拷贝，这样只会在某一个分片上加锁，而不会阻塞整个缓存的读写操作。

Copy-On-Write：

当需要进行快照时，可以创建一个新 map 并将旧数据浅拷贝到新 map 中，在拷贝过程中，读取仍然可以继续，这减少了写锁的时间。

读写锁（sync.RWMutex）：

使用 sync.RWMutex 读写锁来保护数据访问。写操作（如 Add 和 Remove）加写锁，而读取操作（如 Get 和 SnapShot）加读锁。结合分片设计，可以更细粒度地控制锁。
```go
package main

import (
	"fmt"
	"sync"
	"hash/fnv"
)

const ShardCount = 128 // 将 map 分为 128 个分片

type ShardedCache struct {
	shards []*cacheShard
}

type cacheShard struct {
	data map[string]int
	sync.RWMutex
}

// 创建一个新的分片缓存
func NewShardedCache() *ShardedCache {
	shards := make([]*cacheShard, ShardCount)
	for i := 0; i < ShardCount; i++ {
		shards[i] = &cacheShard{
			data: make(map[string]int),
		}
	}
	return &ShardedCache{
		shards: shards,
	}
}

// 哈希函数，根据 key 计算属于哪个分片
func getShard(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % ShardCount
}

// Add 添加键值对
func (sc *ShardedCache) Add(key string, value int) {
	shard := sc.shards[getShard(key)]
	shard.Lock()
	defer shard.Unlock()
	shard.data[key] = value
}

// Remove 删除键值对
func (sc *ShardedCache) Remove(key string) {
	shard := sc.shards[getShard(key)]
	shard.Lock()
	defer shard.Unlock()
	delete(shard.data, key)
}

// Get 获取值
func (sc *ShardedCache) Get(key string) (int, bool) {
	shard := sc.shards[getShard(key)]
	shard.RLock()
	defer shard.RUnlock()
	val, ok := shard.data[key]
	return val, ok
}

// SnapShot 快照
func (sc *ShardedCache) SnapShot() map[string]int {
	snapshot := make(map[string]int)

	var wg sync.WaitGroup
	for _, shard := range sc.shards {
		wg.Add(1)
		go func(s *cacheShard) {
			defer wg.Done()
			s.RLock()
			defer s.RUnlock()
			for k, v := range s.data {
				snapshot[k] = v // 快照复制
			}
		}(shard)
	}
	wg.Wait() // 等待所有分片处理完成
	return snapshot
}

func main() {
	cache := NewShardedCache()

	cache.Add("key1", 100)
	cache.Add("key2", 200)

	val, ok := cache.Get("key1")
	fmt.Println("Get key1:", val, ok)

	snapshot := cache.SnapShot()
	fmt.Println("Snapshot:", snapshot)
}

```
### 总结
优化分析：

分片技术：通过将缓存分片，操作时只锁定某一小部分数据，从而减少锁的粒度，提高并发性。写操作只会锁定其对应的分片，避免整个缓存被锁定。

并发快照：在 SnapShot() 中并发地对各个分片进行浅拷贝，每个分片内部加 RLock 锁确保数据一致性。这样做的好处是，整个 SnapShot() 过程中的锁争用被限制在每个分片内，减少了阻塞时间。

优化方案总结：

分片可以大幅减少锁的粒度，提升并发性能。

Copy-On-Write 适合场景是快照的创建次数较少，而读写操作频繁时可以极大减少锁争用。
通过引入 RWMutex 来区分读锁和写锁，并发性能可以进一步提升。




