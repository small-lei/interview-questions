# leetcode102. 二叉树的层序遍历共享屏幕手撕算法：把一个二叉树做一个层序遍历。（假如说有一个两层的平衡二叉树，要做的事就是把这个二叉树按照广度优先遍历给它遍历出来，对于每一层的数据，第一层按照从左到右遍历，第二层按照从右到左的顺序
```go
type TreeNode struct{
    Valint
    Left*TreeNode
    Right*TreeNode
}
```
#### 要解决二叉树的层序遍历问题，并且要求每一层按顺序交替（从左到右、从右到左）遍历，可以使用广度优先遍历（BFS），并通过一个标志变量来控制每一层的遍历方向。我们可以利用队列来实现层序遍历，同时在遍历每一层时，根据当前层的奇偶性来决定是否反转该层的节点顺序。
```go
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	// 队列实现广度优先搜索
	queue := []*TreeNode{root}
	leftToRight := true // 标记当前层的遍历方向

	for len(queue) > 0 {
		levelSize := len(queue) // 当前层的节点数量
		levelNodes := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// 根据当前层的遍历方向决定节点的顺序
			if leftToRight {
				levelNodes[i] = node.Val
			} else {
				levelNodes[levelSize-1-i] = node.Val
			}

			// 将子节点加入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 添加当前层结果
		result = append(result, levelNodes)

		// 切换下次遍历的方向
		leftToRight = !leftToRight
	}

	return result
}

// 辅助函数用于打印结果
func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	result := zigzagLevelOrder(root)
	for _, level := range result {
		fmt.Println(level)
	}
}

```
#### 代码解释：
数据结构：使用 queue 来实现广度优先遍历。

层序遍历：每次处理一层的所有节点，将下一层的子节点加入队列中。

方向切换：通过 leftToRight 变量标记当前层的遍历方向，当方向为 true 时，从左到右添加节点，否则从右到左。

队列管理：当前层处理完后，将下一层的节点全部加入队列。

#### 输出结果
```text
[1]
[3 2]
[4 5 6 7]
```

# 广度优先和深度优先的区别

#### 广度优先搜索（BFS）和深度优先搜索（DFS）是两种常见的图遍历算法，它们的主要区别在于遍历的顺序和方式。具体区别如下：

#### 遍历策略
广度优先搜索（BFS）：
   从一个起始节点开始，按照层次逐层遍历，即先访问距离起始节点最近的节点，然后再访问更远的节点。
   遍历顺序是"横向"的，优先处理当前层的所有节点，之后再处理下一层。
  
 深度优先搜索（DFS）：
   从起始节点开始，一直深入到某条路径的最深处，直到无法再深入（即到达叶节点或死胡同）后，回溯到上一个节点并尝试另一条路径。
   遍历顺序是"纵向"的，优先探测当前节点的子节点，再返回探测其他未遍历的分支。
####  使用的数据结构
   BFS：
使用 队列（FIFO，先进先出）来存储待访问的节点。
由于按层次遍历，队列用于保持当前层的节点，并依次访问每个节点。

DFS：
使用 栈（可以是显式栈或者递归调用的系统栈，LIFO，后进先出）来存储待访问的节点。
每次深入时，沿着某条路径尽可能走到最深，然后回溯。
#### 应用场景
   BFS：
   更适合用于寻找最短路径问题，因为它是逐层扩展，先到达的节点一定是最短路径上的节点。
   适用于所有节点的同层次遍历，例如：层序遍历二叉树、无权图中求最短路径。
   
DFS：
   更适合用于探索所有可能的路径，例如：找到所有解、全排列、图的连通性检测等。
   在某些场景下，DFS可能会比BFS更节省空间，因为DFS只需跟踪当前路径，而BFS需要保存每一层的所有节点。
#### 时间和空间复杂度
   BFS：
   时间复杂度：O(V + E)，其中 V 是顶点数，E 是边数。
   空间复杂度：最坏情况下为 O(V)，需要存储当前层的所有节点。
   
DFS：
   时间复杂度：O(V + E)。
   空间复杂度：最坏情况下为 O(V)，递归时需要栈空间或显式的栈，栈深度最坏情况下为图的高度（树的深度）。
#### 回溯
   BFS：
   不需要显式的回溯，因为它是逐层遍历的，一层遍历完再进入下一层。
   
DFS：
   需要显式回溯。在某条路径走到尽头时，需要回退到上一个节点，并尝试其他路径。
#### 代码实现
   BFS 代码实现（基于队列）：
```go
func bfs(root *TreeNode) {
    if root == nil {
        return
    }

    queue := []*TreeNode{root}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:] // 出队

        fmt.Println(node.Val)

        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
}

```
#### DFS 代码实现（递归实现）：
```go
func dfs(node *TreeNode) {
    if node == nil {
        return
    }

    fmt.Println(node.Val)
    dfs(node.Left)
    dfs(node.Right)
}
```
#### 总结
BFS 更适合解决层次相关的问题，通常会用在无权图中的最短路径查找、二叉树的层序遍历等。

DFS 更适合解决路径探索、连通性、递归结构等问题，如全排列、拓扑排序、图的连通分量查找等。