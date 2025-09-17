// 数值常量
// 数值常量是高精度的 值

package main
import (
	"fmt"
)

const (
	Big = 1 << 100 // Big 是二进制的1后跟100个0
	Small = Big >> 99 // Small = 2
)

func needInt(x int) int {
	return x*10 +1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main(){
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}