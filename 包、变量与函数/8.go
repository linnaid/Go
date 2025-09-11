// 变量的初始化

package main

import "fmt"

// 必须写上每个变量的值，不能省略任何一个
var i, j int = 1, 2
var b int

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java, b)
}