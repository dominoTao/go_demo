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


func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1,2,3,4))
	t.Log(sum(1,2,3,4,5))
}
// defer 关键字  推迟执行
func Clear() {
	fmt.Println("clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("start ...")
	time.Sleep(time.Second * 3)
	panic("err")  // 报错
	fmt.Println("end")
}