# 用户金币数实时排行榜，如果让你在内存中实现，你将如何设计数据结构？

为了在内存中实现用户金币数的实时排行榜，设计一个高效的数据结构是关键。我建议使用以下组合：

1. Sorted Map（有序映射）
   数据结构: 使用一种可以根据金币数排序的数据结构，比如**红黑树**或其他支持排序的平衡二叉树。
   用途: Sorted Map中的每个键值对为用户ID和金币数。键为金币数，值为用户ID列表（多个用户可能有相同的金币数）。这保证了我们可以快速获取排名、插入和更新。
2. Hash Map（哈希表）
   数据结构: Hash Map用来将用户ID映射到用户当前的金币数。
   用途: 便于快速查找和更新用户的金币数。当某个用户金币数发生变化时，可以通过Hash Map快速查到该用户的当前金币，然后在Sorted Map中更新其位置。
3. 伪代码设计
   go
   复制代码
```go

   // 假设使用Golang
   type Leaderboard struct {
       users      map[string]int        // 用户ID到金币数的映射
       sortedRank *TreeMap[int][]string // 金币数到用户ID列表的映射
   }

    func (lb *Leaderboard) AddOrUpdateUser(userID string, coins int) {
        if oldCoins, exists := lb.users[userID]; exists {
        // 先从 sortedRank 中移除用户
            lb.removeUserFromSortedRank(userID, oldCoins)
        }
        // 更新用户金币数
        lb.users[userID] = coins
        // 将用户插入 sortedRank
        lb.addUserToSortedRank(userID, coins)
    }

    func (lb *Leaderboard) removeUserFromSortedRank(userID string, coins int) {
        if userList, exists := lb.sortedRank[coins]; exists {
            // 移除该用户
            lb.sortedRank[coins] = removeFromList(userList, userID)
            // 如果此金币数没有其他用户，删除该金币数条目
            if len(lb.sortedRank[coins]) == 0 {
                delete(lb.sortedRank, coins)
            }
        }
    }

    func (lb *Leaderboard) addUserToSortedRank(userID string, coins int) {
        if _, exists := lb.sortedRank[coins]; !exists {
            lb.sortedRank[coins] = []string{}
        }
        lb.sortedRank[coins] = append(lb.sortedRank[coins], userID)
    }

	func (lb *Leaderboard) GetTopN(n int) []string {
        var result []string
        for coins, userList := range lb.sortedRank.ReverseOrder() {
            for _, userID := range userList {
                result = append(result, userID)
                if len(result) == n {
                    return result
                }
            }
        }
        return result
    }

```
4. 操作复杂度
   插入/更新: O(log n)，插入或更新用户金币数时，需要从Sorted Map中找到合适的位置。
   获取Top N: O(N)，按排名顺序遍历Sorted Map中的前N个用户。
5. 扩展性考虑
   并发处理: 使用读写锁（RWMutex）来保证并发读写的安全性。
   数据持久化: 可以定期将排行榜数据持久化到磁盘，或者使用Redis等内存数据库来增强扩展性。
   这个设计能够在内存中高效地实现实时更新和读取排行榜。