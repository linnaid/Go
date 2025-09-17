// 结构体

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// 结构体字面量
var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予零值
	v3 = Vertex{}      // X:0 Y:0
	l  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	// 通过 . 访问
	v.X = 4
	fmt.Println(v.X, v.Y)
	// 结构体指针
	p := &v 
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, l, v2, v3)
}