package extension

import (
	"fmt"
	"testing"
)

type Pet struct {

}
func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(name string) {
	p.Speak()
	fmt.Print(" ", name)
}

type Dog struct {
	//p *Pet
	Pet
}

func (d *Dog) Speak() {
	//d.p.Speak()
	fmt.Print("wang")
}
func (d *Dog) SpeakTo(name string)  {
	//d.p.SpeakTo(name)
	d.SpeakTo(name)
}
func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("chao")
	//dog.SpeakTo("chao")
}