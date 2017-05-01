package main

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





