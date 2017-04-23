package main

import "fmt"

func main() {
	data := [3]int{10,20,30}

	for i,x := range data{
		if i== 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}

		fmt.Printf("x:%d,data:%d\n",x,data[i])
	}

	//使用切片来替换数组 复制的是slice自身
	for i,x := range data[:]{
		if i== 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}

		fmt.Printf("x:%d,data:%d\n",x,data[i])
	}
}
