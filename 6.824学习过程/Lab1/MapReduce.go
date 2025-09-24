
package main

import (
	"fmt"
)

// Map只记录哪些字符出现了,例如{你,1}{w,1}{你,1}只记录字符
func Map(s string) [][2]interface{} {
	word_num := make([][2]interface{}, 0)
	for _, c := range s {
		word_num = append(word_num, [2]interface{}{c, 1})
	}
	return word_num
}

// Shuffle 记录每个字符的数量，例如：word_nums[2] = {1,1},代表它出现了两次
func Shuffle(pairs [][2]interface{}) map[rune][]int {
	word_nums := make(map[rune][]int)
	for _, pair := range pairs {
		ch := pair[0].(rune)
		num := pair[1].(int)
		word_nums[ch] = append(word_nums[ch], num)
	}
	return word_nums
}

// 计算字符的数量，得到最终答案
func Reduce(word_nums map[rune][]int) map[rune]int {
	ans := make(map[rune]int)
	for ch, word_num := range word_nums {
		sum := 0
		for _, a := range word_num {
			sum += a
		}
		ans[ch] = sum
	}
	return ans
}

func main() {
	var s string
	fmt.Printf("请输入一个字符串：")
	fmt.Scanln(&s)

	pairs := Map(s)
	word_nums := Shuffle(pairs)
	result := Reduce(word_nums)

	fmt.Println("字符出现的个数：")
	for c, num := range result {
		fmt.Printf("%c : %d\n", c, num)
	}
}