package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}

	dummy := &ListNode{Next: head}
	pre := dummy

	// 移动到 m 的前一个节点
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}

	// 开始反转
	cur := pre.Next            // 1. cur指向需要反转的第一个节点
	for i := 0; i < n-m; i++ { // 2. 循环n-m次，反转指定区间
		next := cur.Next     // 3. 保存当前节点的下一个节点
		cur.Next = next.Next // 4. 当前节点指向下下个节点(跳过next)
		next.Next = pre.Next // 5. next节点指向pre的下一个节点(插入到前面)
		pre.Next = next      // 6. pre的下一个节点更新为next
	}

	return dummy.Next
}

// 辅助函数：创建链表
func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return dummy.Next
}

// 辅助函数：打印链表
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// 测试用例
	head := createList([]int{1, 2, 3, 4, 5})
	fmt.Print("原链表: ")
	printList(head)

	reversed := reverseBetween(head, 2, 4)
	fmt.Print("反转后: ")
	printList(reversed)
}
