// 函数1

package main

import "fmt"

// 注意；类型在变量名后
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 2));
	fmt.Println(add2(1, 2));
}


// 函数2

// 连续两个已命名形参类型相同可以省略前面的类型声明，只保留最后一个
func add2(x, y int) int {
	return x + y
}

