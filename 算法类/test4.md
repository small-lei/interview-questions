# 代码题，最长公共子串
求两个字符串的最长公共子串，可以使用动态规划（Dynamic Programming，简称 DP）来解决。动态规划的思路是用一个二维数组来记录子串匹配的情况，逐步构建出最长的公共子串。

动态规划解法：
```text
创建一个二维数组 dp，其中 dp[i][j] 表示字符串 s1 的前 i 个字符与字符串 s2 的前 j 个字符的最长公共后缀长度。

如果 s1[i-1] == s2[j-1]，则 dp[i][j] = dp[i-1][j-1] + 1。

如果不相等，则 dp[i][j] = 0。
通过遍历整个二维数组，记录最大公共后缀的长度，并最终返回最长公共子串。
```
### 示例代码
```go
package main

import (
	"fmt"
)

// longestCommonSubstring 求解两个字符串的最长公共子串
func longestCommonSubstring(s1, s2 string) string {
	m, n := len(s1), len(s2)
	if m == 0 || n == 0 {
		return ""
	}

	// dp数组，记录每个字符配对时的最长公共后缀长度
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 记录最长公共子串的长度及结束位置
	longestLength := 0
	endIndex := 0

	// 填充dp数组
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > longestLength {
					longestLength = dp[i][j]
					endIndex = i // 记录子串结束的位置
				}
			} else {
				dp[i][j] = 0
			}
		}
	}

	// 返回最长公共子串
	return s1[endIndex-longestLength : endIndex]
}

func main() {
	s1 := "abcdfg"
	s2 := "abdfg"
	result := longestCommonSubstring(s1, s2)
	fmt.Printf("最长公共子串是: %s\n", result)
}

```

### 代码解析：
二维数组 dp：
```text
dp[i][j] 表示字符串 s1 的前 i 个字符和字符串 s2 的前 j 个字符的最长公共后缀长度。
当 s1[i-1] == s2[j-1] 时，说明在当前字符处有匹配，因此 dp[i][j] = dp[i-1][j-1] + 1。
当 s1[i-1] != s2[j-1] 时，表示当前字符不匹配，公共后缀长度为 0。
最大子串的记录：

在每次找到更长的公共子串时，更新 longestLength 和 endIndex，其中 endIndex 是 s1 中公共子串的结束位置。
最终通过 endIndex 和 longestLength 从 s1 中截取最长公共子串。
时间复杂度：
时间复杂度是 O(m * n)，其中 m 是字符串 s1 的长度，n 是字符串 s2 的长度。
```