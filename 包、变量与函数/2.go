// 导入

package main

// 这是分组导入，也可以分别导入
import (
	"fmt" // 用于打印输出
	"math" 
)
func main() {
	fmt.Printf("现在你有了 %g 个问题.\n", math.Sqrt(7))
}

