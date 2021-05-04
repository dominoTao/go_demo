package src_ch3_type_test

import "testing"
type MyInt int64

const (
	Readable = 1 << iota
	Writable
	Executable
)
func TestCode(t *testing.T) {
	var s string
	//s = "hello"
	//t.Log(len(s))
	////s[1] = '3'
	s = "\xE4\xB8\xAD"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

}

func TestBitClear(t *testing.T)  {
	a := 7  // 0111
	a = a &^ Readable  // 清除可读
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

func TestImplicit(t *testing.T){
	var a int32 = 1
	var b int64
	b = int64(a)

	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}
func TestPoint(t *testing.T) {
	a:=1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}
func TestString(t *testing.T) {
	var s string
	t.Log("*"+s+"*")
	t.Log(len(s))
}
func TestS(t *testing.T) {
	a := [...]int{1,2,3,4}
	b := [...]int{1,2,3,5}
	//c := [...]int{1,2,3,4, 3}
	d := [...]int{1,2,3,4}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}