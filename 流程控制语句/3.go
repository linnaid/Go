// 练习 循环与函数

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		pos := z
		z -= ((z * z) - x) / (2 * z);
		if math.Abs(z - pos) < 1e-10 {
			break
		} 
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
