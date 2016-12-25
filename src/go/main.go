package main

import (
	"errors"
	"fmt"
)

func main() {
	// var x int32 //0
	// var s = "hello world"
	// a := 20       //定义一定要被使用
	// b := "asdasd" //字符串必须使用双引号
	// c := 100

	// if c > 0 {
	// 	fmt.Println(c)
	// } else if c < 0 {
	// 	fmt.Println(c + 1)
	// } else {
	// 	fmt.Println("0")
	// }

	// fmt.Println(x, s, a)

	x := []int{100, 101, 102}
	//i 表示索引之爱 n是从数组中取出来的
	for i, n := range x {
		fmt.Println(i, ":", n)
	}

	a, b := 10, 2
	c, err := div(a, b)

	fmt.Println(c, err)
}

//后面的 (int,errors) 是返回值类型
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

// func circulation() {
// 	switch {
// 	case x > 0:
// 		println("x")
// 	case x < 0:
// 		println("x")
// 	default:
// 		println("x")
// 	}
// }
