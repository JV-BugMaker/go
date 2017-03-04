package main

import "fmt"

func main() {
	var a struct{
		x int `x` //struct tag
		s string `s`
	}

	var b struct{
		x int
		s string
	}
	b = a  //错误 忽略了struct tag也是struct 签名的一部分

	fmt.Println(b)

	var c func(int,string)
	var d func(string,int)

	d = c //错误 函数参数的顺序也是需要考虑的

	d("s",1)
	//未命名类型转换规则:
	//所属类型相同
	//基础类型相同，且其中一个是未命名函数类型
	//数据类型相同，将双向通道赋值给单向通道，且其中一个为未命名类型
	//将默认值nil赋值给切片、字典、通道、指针、函数或接口
	//对象实现了目标接口
	type data [2]int
	var e data = [2]int{1,2} //基础类型相同，右值为未命名类型
	fmt.Println(e)
	k := make(chan int,2)
	var g chan<- int = k //双向通道转换为单向通道，其中g为未命名类型
	g <- 2

}