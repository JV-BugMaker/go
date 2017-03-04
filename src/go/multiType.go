package main

import "fmt"

func main() {
	//定义组
	type (
		//结构体
		user struct {
			name string
			age uint8
		}
		event func(string) bool
	)

	u := user{"JV",24}
	fmt.Println(u)
	
	var f event = func(s string) bool {
		println(s)
		return s !=""
	}

	f("abc")
}
