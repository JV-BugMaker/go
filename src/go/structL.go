package main

import (
	"fmt"
	"reflect"
)

//struct是个符合类型 字段名必须是唯一 可用"_"补位 支持使用自身指针类型成员

//只有在所有字段全部支持时，才可做相等操作

//空结构可作为通道元素类型 用于事件通知

type attr struct {
	perm int
}

type file struct {
	name string
	attr  //只有类型名 没有字段名
}

func main() {
	exit := make(chan struct{})

	go func() {
		println("hello world")
		exit <- struct{}{}
	}()

	<-exit
	println("end.")
}

func pointer(){
	type data struct {
		*int
		string
	}

	x := 100
	d := data{
		int:&x,
		string:"abc",
	}

	fmt.Println("%#v\n",d)

}

//不能同时嵌入基础类型和其指针基础类型
//type data struct {
//	*int
//	int
//}

//字段标签不是注释 而是用来对字段进行描述的元数据
//在运行期间，可用反射获取标签信息 ，它常被用作格式校验 数据库关系映射

type user struct {
	name string `昵称`
	sex byte `性别`
}

func tag(){
	u := user{"Tom",1}
	v := reflect.ValueOf(u)
	t := v.Type()

	for i,n := 0,t.NumField();i<n ;i++  {
		fmt.Printf("%s:%v\n",t.Field(i).Tag,v.Field(i))
	}
}
//在分配内存的时候 字段必须做对齐处理 通常以所有字段中最长的基础类型宽度为标准