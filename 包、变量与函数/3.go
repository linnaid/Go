// 导出名

package main

import(
	"fmt"
	"math"
)

// pi是未导出名(只能在包内使用)，所以不能使用
// Pi是已导出名，所以可以使用
func main() {
	fmt.Println(math.Pi)
}

// 类的私有变量