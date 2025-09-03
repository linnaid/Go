// 包
// go不需要显式插入分号，go语言会自动在句末加入

package main

import (
	"fmt"
	"math/rand" // 用于生成随机数
)
func main() {
	fmt.Println("我最喜欢的数字是 ", rand.Intn(10))
}