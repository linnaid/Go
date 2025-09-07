// 带名字的返回值

package main

import "fmt"


func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum / x
	// 没有参数的 return 语句会直接返回已命名的返回值，也就是「裸」返回值
	return 
}

func main() {
	fmt.Println(split(17))
}