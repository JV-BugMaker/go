package main

import "fmt"

//可变参数其实就是切片
func testParams(a ...int){
	for i:= range a {
		a[i] += 10
	}
}
func main() {
	a := []int{100,200,300}
	testParams(a...)

	fmt.Println(a)
}