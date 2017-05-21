package main

import (
	"reflect"
	"fmt"
)
//只有在获取结构体指针的基类型后 才能遍历它的字段
type user struct {
	name string
	age int
}

type manager struct {
	user   //匿名字段
	title string
}

func main() {
	var m manager
	t := reflect.TypeOf(&m)

	if t.Kind() == reflect.Ptr {    //获取指针的基类型
		t = t.Elem()
	}

	for i:=0;i < t.NumField();i++{
		f := t.Field(i)
		fmt.Println(f.Name,f.Type,f.Offset)
		if f.Anonymous {   //输出匿名字段结构
			for x := 0;x < f.Type.NumField();x++{
				af := f.Type.Field(x)
				fmt.Println(" ",af.Name,af.Type)
			}
		}
	}
	//user main.user 0
	//name string
	//age int
	//title string 24

}
