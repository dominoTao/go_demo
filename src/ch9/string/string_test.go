package s_test

import "testing"

func TestString(t *testing.T) {
	s := "中"
	t.Log(len(s)) // byte 切片中  的长度  \xE4  \xB8  \xAD
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for e, c := range s {
		t.Logf("%[1]c %[1]x", c)
		t.Logf("%[1]c %[1]d", c)
		t.Log(e)
	}
}