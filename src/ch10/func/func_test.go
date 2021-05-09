package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)
// 函数  一等公民
func returnMultiValues() (int, int)  {
	return rand.Intn(10), rand.Intn(20)
}

func slowFun(op int) int {
	time.Sleep(time.Second * 3)
	return op
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func TestFn(t *testing.T) {
	//a, b := returnMultiValues()
	//t.Log(a, b)
	//c, _ := returnMultiValues()
	//t.Log(c)

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))

}