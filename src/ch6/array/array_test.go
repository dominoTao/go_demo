package array

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(arr[1], arr[2], arr[0])

	arr1 := [4]int{1,2,3,4}
	t.Log(arr1)
	t.Log(arr1[1])

	arr2 := [...]int{1,2,3,4,5,6,7,8}
	arr2[1] = 0
	t.Log(arr2)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1,3,4,5}
	//for i:=0; i< len(arr3); i++ {
	//	t.Log(arr3[i])
	//}
	// 新写法  下划线表示不关心索引值
	for _, e := range arr3 {
		//t.Log(_, e)
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1,2,3,4,5}
	arr3_sec := arr3[:3]
	t.Log(arr3_sec)
	arr3_sec1 := arr3[3:]
	t.Log(arr3_sec1)
	arr3_sec2 := arr3[:]
	t.Log(arr3_sec2)
}