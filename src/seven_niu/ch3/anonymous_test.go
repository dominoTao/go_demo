package anonymous_combination_test

import (
	"log"
	"testing"
)

type Base struct {
	Name string
}

func (base *Base) Foo()  {
	log.Println("func Foo from Base")
}

func (base *Base) Bar() {
	log.Println("func Bar from Base")
}

type Foo struct {
	Base
}

func (foo *Foo) Bar() {
	log.Println("func Bar from Foo")
}

func TestExtents(t *testing.T){
	b := new(Base)
	b.Foo()
	b.Bar()
	f := &Foo{}
	f.Foo()
	f.Bar()
}



















type Job struct {
	Command string
	*log.Logger
}
func (job *Job) Start(){
	job.Println("starting now...")
	job.Println("started")
}

func TestJob(t *testing.T) {
	// error
	var i *log.Logger
	i = new(log.Logger)
	j := &Job{"command demo", i}
	j.Start()
}




type X struct {
	Name string
}
type Y struct {
	X
	Name string
}

type Logger struct {
	Level int
}
type YY struct {
	*Logger
	Name string
	// 匿名组合类 名称相当于去除包名  所以相当于定义了两个Logger  虽然类型不一样  但会编译报错
	//*log.Logger
}