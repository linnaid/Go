// 常量
// 常量可以是字符、字符串、布尔值或数值

package main

import "fmt"

// 常量不能用 := 语法声明
const Pi = 3.14

func main() {
	const World = "world";
	fmt.Println("Hello", World);
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}