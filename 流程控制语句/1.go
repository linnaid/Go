// for 循环
// for 在 go 中也被当作 while 使用

package main

import "fmt"

func main() {
	sum := 0
	for  i := 0; i < 10; i++ {
		sum += i;
	}
	fmt.Println(sum);
	// while
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum);
	// 无限循环
	for{
	}
}