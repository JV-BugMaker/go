package main

import "fmt"

//引用类型 特指slice、map、channel这三种预定义类型
//引用类型则必须make函数创建，编译器会将make转换为目标类型专用的创建函数，以确保完成全部内存和相关属性初始化

//返回一个整型数组
func mkslice() []int{
	s := make([]int,0,10)
	s = append(s,100)
	return s
}

//返回一个key为string 值为整型的map
func mkmap() map[string] int{
	m := make(map[string]int)
	m["a"] = 1
	return m
}

func main() {
	m := mkmap()
	println(m["a"])

	s := mkslice()
	println(s[0])
	//new函数也可为引用类型分配内存 但这是不完整创建。以字典为列，它仅分配字典类型本身（实际就是个指针包装）所需内存
	//并没有分配键值存储内存
	p := new(map[string]int)  //函数new返回指针
	m:= *p
	m["a"] = 1 //panic:assigment to entry in nil map
	fmt.Println(m)
}