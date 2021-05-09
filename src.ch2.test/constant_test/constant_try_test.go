package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday

)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTestTry(t *testing.T)  {
	t.Log(Monday, Tuesday, Wednesday)
	t.Log(Readable, Writable, Executable)
	a:=3 // 011
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}