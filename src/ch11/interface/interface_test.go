package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {

}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello world\")";
}

func TestClient(t *testing.T) {

	var p Programmer
	p = new(GoProgrammer)

	t.Log(p.WriteHelloWorld())

	var prog Programmer = &GoProgrammer{}
	prog.WriteHelloWorld()
}