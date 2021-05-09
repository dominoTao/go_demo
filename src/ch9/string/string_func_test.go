package s_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringfn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts{
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}


func TestConv(t *testing.T) {
	s:=strconv.Itoa(10)
	t.Log("str" + s)
	if i, err:=strconv.Atoi("102dd"); err==nil{
		t.Log(10+i)

	}else {
		t.Log(err)
	}
}