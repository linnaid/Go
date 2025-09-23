// switch 分支

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go 运行的系统环境: ")
	switch os := runtime.GOOS; os {
	// switch 的 case 无需为常量，且取值不限于整数
	case "darwin":
		fmt.Println("macOS.")
	//  Go 只会运行选定的 case，而非之后所有的 case
	// 相当于这些语言中为每个 case 后面自动添加了所需的 break 语句
	case "linux":
		fmt.Println("Linux.");
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// switch 的求值顺序
	fmt.Println("周六是哪天？")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("今天。")
	case today + 1:
		fmt.Println("明天。")
	case today + 2:
		fmt.Println("后天。")
	default:
		fmt.Println("很多天后。")
	}

	// 无条件 switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("早上好！")
	case t.Hour() < 17:
		fmt.Println("下午好！")
	default:
		fmt.Println("晚上好！")
	}
}




