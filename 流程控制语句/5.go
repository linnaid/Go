// defer 推迟

package main

import (
	"fmt"
)

func main() {
	// defer 语句会将函数推迟到外层函数返回之后执行
	defer fmt.Println("world")

	// 推迟调用的函数调用会被压入一个栈中。 当外层函数返回时，被推迟的调用会按照后进先出的顺序调用
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("hello")
}