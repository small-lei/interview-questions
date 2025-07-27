# 算法题：最长有效括号（LeetCode 32) ?
题目 LeetCode 32: 最长有效括号 要求我们找到字符串中最长的包含有效括号的子串的长度。有效括号指的是配对正确的 ()。

栈
使用栈来存储左括号的索引。当遇到右括号时，如果栈不为空，说明遇到了一对有效的括号，可以弹出栈顶元素。
通过栈中的元素计算最长的有效括号子串长度。
```go
package main

import "fmt"

func longestValidParenthesesStack(s string) int {
    stack := []int{-1} // 初始化栈，-1作为一个哨兵
    maxLen := 0
    
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack)-1] // 弹出栈顶
            if len(stack) == 0 {
                // 如果栈为空，意味着无法匹配当前右括号，记录当前索引
                stack = append(stack, i)
            } else {
                // 否则计算当前有效括号的长度
                maxLen = max(maxLen, i - stack[len(stack)-1])
            }
        }
    }
    
    return maxLen
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    s := "(()))())("
    fmt.Println("Longest valid parentheses length:", longestValidParenthesesStack(s)) // 输出: 4
}
```
### 复杂度分析：
```text
时间复杂度：O(n)，每个括号都只会被压栈和弹栈一次。
空间复杂度：O(n)，栈的大小最多为 n。
```