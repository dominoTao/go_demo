package map_ext

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int)int{}
	m[1] = func(op int) int {return op}
	m[2] = func(op int) int {return op * op}
	m[3] = func(op int) int {return op * op * op}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n]{
		t.Logf("mySet[%d] is existing , mySet[%d] = ", n, n)
		t.Log(mySet[n])
	} else {
		t.Logf("mySet[%d] is not existing.", n)
	}
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 1)
	m := 1
	if mySet[m]{
		t.Logf("mySet[%d] is existing , mySet[%d] = ", m, m)
		t.Log(mySet[m])
	} else {
		t.Logf("mySet[%d] is not existing.", m)
	}
	t.Log(len(mySet))

}