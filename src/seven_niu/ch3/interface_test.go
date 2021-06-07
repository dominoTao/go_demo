package anonymous_combination_test

import (
	"log"
	"testing"
)


type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}
type IReader interface {
	Read(buf []byte) (n int, err error)
}
type IWriter interface {
	Write(buf []byte) (n int, err error)
}
type IClose interface {
	Close() error
}


type File struct {
	// 虽然该类没有继承上面的接口，但是下面方法已经实现了接口（实现了接口的方法），就可以进行赋值，如测试方法TestStruct
}
func (f *File) Read(buf []byte) (n int, err error){
	return
}
func (f *File) Write(buf []byte) (n int, err error){
	return
}
func (f *File) Seek(off int64, whence int) (pos int64, err error){
	return
}
func (f *File) Close() error{
	return nil
}

func TestStruct(t *testing.T) {
	var file1 IClose= &File{}
	var file2 IReader = &File{}
	var file3 IWriter = new(File)
	var file4 IFile = new(File)
	file1.Close()
	file2.Read(nil)
	file3.Write(nil)
	file4.Seek(0, 0)
}
























/**
实例赋值给接口

接口赋值给接口
 */

type Integer int
/*
Less 方法可以自动生成 Less1 方法,这样一来  *Integer 就即存在Less方法 又存在Add方法， 满足LessAdder接口，
而Add 方法并不能生成 Add1方法，因为(&a).Add(b)只改变参数a ，对外部对象没有影响，不符合预期
不满足LessAdder接口所以 接口回调的时候(即 将Integer对象赋值给LessAdder接口的时候) 必须在对象前加上&  如测试方法TestInterface
 */
func TestInterface(t *testing.T){
	var a Integer = 1
	log.Print(a)
	//var b LessAdder = a
	//log.Print(b)
	var c LessAdder = &a
	log.Print(c)
}
func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Less1(b Integer) bool {

	return (*a).Less(b)
}
func (a *Integer) Add(b Integer) {
	*a += b
}

func (a Integer) Add1(b Integer) {
	(&a).Add(b)
}




type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}









type ReadWriter1 interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}
type ReadWriter2 interface {
	Write(buf []byte) (n int, err error)
	Read(buf []byte) (n int, err error)
}
func TestSameInterface(t *testing.T) {
	/*
	ReadWriter1和ReadWriter2是同一个接口  可以相互赋值
	 */
	var file1 ReadWriter1 = new(File)
	var file2 ReadWriter2 = file1
	var file3 ReadWriter1 = file2
	log.Println(file3)

	// 这里会报错 因为 file4没有 IFile中别的方法，  不过可以通过接口查询语法 做到
	var file4 IReader = new(File)
	//var file5 IFile = file4
	// 接口查询
	if filee, ok := file4.(*File); ok {  // 询问接口file4 ，指向的对象是不是*File类型
		log.Println(filee)
	}

	var v1 interface{}
	switch v:=v1.(type) {
	case int:log.Println(int(v))
	case string:log.Println(string(v))
	default:
		log.Println("error type change")
	}

	var file6 IFile = &File{}
	var file7 IReader = file6
	log.Println(file7)
}


type Stringer interface {
	String() string
}
func Println(args ...interface{}) {
	for _, arg := range args {
		switch v := arg.(type) {
		case int:log.Println(int(v))
		case string:log.Println(string(v))
		default:
			if v, ok := arg.(Stringer); ok {
				val := v.String()
				log.Println(val)
			} else {

			}
		}
	}
}