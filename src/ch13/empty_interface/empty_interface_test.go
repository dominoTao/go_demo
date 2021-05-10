package empty_interface

import (
	"fmt"
	"os"
	"testing"
)
/**
空接口
 */
func DoSomething(p interface{})  {
	//if i, ok := p.(int); ok{
	//	fmt.Println("Integer  ", i)
	//	return
	//}
	//if s,ok:=p.(string);ok{
	//	fmt.Println("string  ", s)
	//	return
	//}
	//fmt.Println("Unknown type.")

	switch v:= p.(type) { // 断言
	case int:
		fmt.Println("Integer ", v)
	case string:
		fmt.Println("String ", v)
	default:
		fmt.Println("Unknown type.")
	}

}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething(os.TempDir())
}