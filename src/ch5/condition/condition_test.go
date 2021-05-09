package condition

import (
	"fmt"
	"runtime"
	"testing"
)

func TestIfMultiSec(t *testing.T) {
	// if 声明; 条件 {}
	if a := 1 == 1; a {
		t.Log("1 == 1")
	}

	switch os:=runtime.GOOS; os {
	case "darwin",os:  // 可以为多个值  并且 case匹配到  不会走defaul
		t.Log("os x")
	case "linux":
		t.Log("linux")
	default:
		fmt.Printf("%s", os)
		fmt.Println()
	}
	// 第二种用法  类似多个if语句
	num := 1
	switch  {
	case num < 0:
		fmt.Println("<0")
	case num >= 0 && num < 8:
		fmt.Println("0-8")
	default:
		fmt.Println("+++")
	}

	for i:= 0; i < 5; i++ {
		switch  {
		case i % 2 == 0:
			t.Log("Even")
		case i % 2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknown")
		}
	}

	//if v, err := somefun(); err == nil {
	//	t.Log("true")
	//} else {
	//	t.Log("false")
	//}
}
//func somefun() [...]int {
//	return [...]int{1}
//}