package operator_test

import "testing"

/**
按位置零运算符    &^
右边为1  无论左边是0还是1  都会把左边的二进制位清零
右边为0  则原来左边的是什么  结果就是什么
1 &^ 0 == 1
1 &^ 1 == 0
0 &^ 1 == 0
0 &^ 0 == 0
 */
const (
	Readable = 1 << iota    // 001
	Writable				// 010
	Executable				// 100
)
func TestConstantTestTry(t *testing.T)  {
	a:=3 // 011
	a = a &^ Readable   // 011 &^ 001  == 010
	a = a &^ Writable   // 010 &^ 010  == 000
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
func TestCCC(t *testing.T) {
	t.Log(1 &^ 0)
	t.Log(1 &^ 1)
	t.Log(0 &^ 1)
	t.Log(0 &^ 0)
}

func TestCompareArray(t *testing.T)  {

	// 元素个数 顺序不同 都是false
	a := [...]int{1,2,3,4}
	b := [...]int{1,3,4,2}
	//c := [...]int{1,2,3,4,5}
	d := [...]int{1,2,3,4}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}