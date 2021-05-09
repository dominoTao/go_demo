package main
import (
	"fmt"
)
type Speaker interface {
	Hello()
}
type User struct {
	Name string
	Age int
}
func (this *User) Hello() {
	fmt.Println("hello my name is ", this.Name)
}

func main()  {
	var s Speaker
	//s = User{"wss", 10}
	s = &User{"wss", 10}
	s.Hello()
}