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


	//类型判断上 必须选择正确的方式
	//使用kind进行判断

}

//传入对象应该区分基类型和指针类型 它们不属于同一类型

func main2() {
	x := 100

	tx,tp := reflect.TypeOf(x),reflect.TypeOf(&x)  //传入基本类型 和  指针类型

	fmt.Println(tx,tp,tx == tp)
	fmt.Println(tx.Kind(),tp.Kind())
	fmt.Println(tx == tp.Elem())
	//方法Elem 返回指针、数组、切片、字典（值）或通道的基类型
}

