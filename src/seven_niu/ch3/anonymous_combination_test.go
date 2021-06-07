package anonymous_combination_test

import (
	"log"
	"testing"
)
/**
结构体
 */
type Rect struct {
	X, Y, Width, Height float64
}

func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

func (r *Rect) Area() float64 {
	return r.Width * r.Height
}


func TestAC(t *testing.T) {
	rect1 := new(Rect)
	log.Println(rect1)
	rect2 := &Rect{}
	log.Println(rect2)
	rect3 := &Rect{0, 0, 100, 100}
	log.Println(rect3)
	rect4 := &Rect{Width: 20, Height: 30}
	log.Println(rect4)
	rect := NewRect(2, 3, 4, 5)
	log.Println(rect)
	area := rect.Area()
	log.Println(area)
}