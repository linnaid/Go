// 多返回值

package main

import "fmt"

// 函数可以返回任意数量的返回值，如下我设置了两个返回值
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// := 声明并初始化变量(短变量声明运算符)
	// 如果已经声明过，则重新赋值
	// 不能在函数外使用
	a, b := swap("nihao", "linna")
	fmt.Println(a, b)
}