package main

import "unsafe"

func main() {
	//定义一个常量块
	//const (
	//	ptrSize = unsafe.Sizeof(uintptr(0))
	//	strSize = len("hello world")
	//	x,y int = 99,-99
	//	b byte = byte(x)
	//	n = uint(y)
	//)

	const (
		x uint16 = 120
		y	//与上一行x类型、右值相同
		s = "abc"
	)



}
