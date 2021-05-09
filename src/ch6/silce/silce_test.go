package silce

import "testing"
// 切片
func TestSliceInit(t *testing.T)  {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))
	s1 := []int{1,2,3,4}
	t.Log(len(s1),  cap(s1))


	s2 := make([]int, 3, 5)


	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2]/*, s2[3]*/)
	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i:=0; i<10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year:=[]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
// 修改了summer的值  Q2的值也会被修改
	summer[0] = "Unknown"
	t.Log(Q2)
	t.Log(summer)
}

func TestSliceCompare(t *testing.T) {
	// 定义两个切片
	//a:=[]int{1,2,3,4}
	//b:=[]int{1,2,3,4}
	// invalid operation: a == b (slice can only be compared to nil)
	//if a==b {
	//	t.Log("equal")
	//}
}