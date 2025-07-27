# 写个小代码假如说现在有1w个字符串，让你来拼接，你会怎么实现？
要拼接大量字符串，例如有 1 万个字符串时，建议使用**strings.Builder**来提高效率。直接使用 + 或 += 操作符拼接字符串，可能会频繁地分配内存并复制内容，效率较低。而 strings.Builder 是 Go 官方推荐的高效方式，用于拼接多个字符串。

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 假设我们有 1 万个字符串
	stringsList := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		stringsList[i] = fmt.Sprintf("str%d", i)
	}

	// 使用 strings.Builder 来高效拼接字符串
	var builder strings.Builder

	for _, str := range stringsList {
		builder.WriteString(str)
	}

	// 输出拼接后的字符串长度，避免打印过多内容
	result := builder.String()
	fmt.Printf("拼接后的字符串长度为: %d\n", len(result))
}
```

### 代码解析：
```text
初始化字符串列表：创建了 1 万个字符串，存储在 stringsList 切片中。
strings.Builder：使用 strings.Builder 来高效地拼接字符串，避免频繁的内存分配和复制。
builder.WriteString(str)：逐个将字符串写入 builder。
生成最终结果：通过 builder.String() 获取拼接后的完整字符串。
使用 strings.Builder 的好处：
减少内存分配：相比直接用 + 号拼接，Builder 内部维护了一个动态扩容的缓冲区，减少内存的重新分配。
提升性能：特别是在拼接大量字符串时性能优势明显。
时间复杂度：
时间复杂度为 O(n)，其中 n 是所有字符串的总长度。
```
