package main

import (
	"fmt"
	"reflect"
)
//反射能够股让我们探知对象的类型信息和内存结构
//反射操作需要的全部信息都源自接口变量。接口变量除存储自身类型外，还会保存实际对象的类型数据


//func TypeOf(i interface{}) Type

//func ValueOf(i interface{}) Value
//反射入口函数 会将任何传入的对象转换为接口类型

type X int

func main() {
	var a X = 100

	t := reflect.TypeOf(a)

	fmt.Println(t.Name(),t.Kind())  //type 表示真实类型  kind表示基础接口类型
}