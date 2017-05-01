package main

import "github.com/rogpeppe/godef/go/ast"

//接口也是一种结构

//不能有字段
//不能定义自己的方法
//只能声明方法 不能实现方法
//可嵌入其他接口类型
//通常以er作为名称后缀

//type tester interface {
//	test()
//	string() string
//}

type data struct {}

func (*data) test(){}
func (data) string() string {return ""}

func pp(a stringer){
	println(a.string())
}

func main() {
	var d data

	var t tester = &d
	t.test()
	println(t.string())
	//==============华丽丽的分割线=============
	var d2 data

	var t2 tester = &d2

	pp(t2)   //隐式转换

	var s stringer = t //超集转换为子集

	println(s.string())
}


type stringer interface {
	string() string
}

type tester interface {
	stringer
	test()
}


//执行机制
//接口使用一个名为itab的结构存储运行期间所需的相关类型信息

type iface struct {
	tab *itab  //类型信息
	data unsafe.Pointer //实际对象指针
}

type itab struct {
	inter *interfaceType //接口类型
	_type *_type //实际对象类型
	fun [1]uintptr //实际对象方法地址
}

//将对象赋值给接口对象变量时，会复制该对象

//我们无法修改接口存储的复制品 因为它是unaddressable的


func main2(){
	//将指针赋值给接口变量
	var t interface{} = &d

	//操作指针 进行修改
	t.(*data).x = 200
}
//只有当itab 和 data 都是nil的时候  接口才是nil




//================华丽丽的分割线


type data int

func (d data) String() string{
	return fmt.Sprintf("data:%d",d)
}

func main() {
	var d data = 15
	var x interface{} = d
	//ok-idiom方式能够避免转换失败不至于进入panic
	if n,ok := x.(fmt.Stringer);ok{
		//转换为更具体的接口
		fmt.Println(n)
	}

	if d2, ok := x.(data);ok{
		//转换回原始类型
		fmt.Println(d2)
	}

	e:=x.(error)
	fmt.Println(e)
}