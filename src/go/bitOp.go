package main

import "fmt"

const (
	read byte = 1 << iota
	write
	exec
	freeze
)

func main() {
	////x := 23
	//////无效操作 必须是无符号整数 或可以转换的无类型常量
	//////b := 1<<x
	////a := 1 << 3
	//
	////var s uint = 3
	////b := 1.0 << s // 错误 因为b并没有提供类型 所以编译器按照1.0是浮点型 浮点型不能进行位操作
	//
	//var b int32 = 1.0 << 3
	//
	//print
	/***************以上是对位操作说明********************/

	a := read | write | freeze
	b := read | freeze | exec
	//按位清楚 它将左右对应二进制都为1的重置为0，以达到一次清除多个标记位的目的
	c := a &^ b

	fmt.Println(c,read)

}