package main

import (
	"reflect"
	"fmt"
	"net/http"
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
	//结构体 指针
	t := reflect.TypeOf(&m)

	if t.Kind() == reflect.Ptr {    //获取指针的基类型
		//找到基类型
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

//对于匿名字段 可用多级索引 直接访问
func main2() {
	var m manager
	t := reflect.TypeOf(m)

	name,_ := t.FieldByName("name") //按名称查找
	fmt.Println(name.Name,name.Type)


	age := t.FieldByIndex([]int{0,1})   //按多级索引查找
	fmt.Println(age.Name,age.Type)

}

//输出方法集时 区分基类型和指针类型

func main3() {
	var s http.Server
	t := reflect.TypeOf(s)

	//反射能探知当前包或外包的非导出结构成员
	for i:=0;i<t.NumField();i++{
		fmt.Println(t.Field(i).Name)
	}
}

//获取struct结构体的tag
//f.Tag.Get

