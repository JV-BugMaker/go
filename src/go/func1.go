package main

import "fmt"

type N int

func (n N) toString() string{
	return fmt.Sprintf("%#x",n)
}

func (N) toString2() string{
	println("hi!")
}

func main() {
	var a N = 25
	println(a.toString())
}

//方法不支持重载
//如方法内部并不引用实例，可省略参数 仅保留类型

//如何选择方法的receiver类型？
//要修改实例参数，用*T
//无须修改状态的小对象活固定值,用T
//大对象建议是*T，以减少复制成本
//引用类型、字符串、函数等指针包装对象，直接用T
//若包含Mutex等同步字段,用*T,避免因赋值造成锁操作无效。
//其他无法确定的情况，都用*T
//指针类型必须是合法的，不能使用多级指针调用方法。