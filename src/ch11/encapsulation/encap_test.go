package encap

import "testing"

type Employee struct {
	Id string
	Name string
	Age int
}

func TestEmployeeEncap(t *testing.T)  {
	em := &Employee{}
	t.Log(em)
	em1 := new(Employee)
	t.Log(em1)
}
