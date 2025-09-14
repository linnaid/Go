// 类型转换

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// 类型转换
	var z uint = uint(f)
	fmt.Println(x, y, z)
}