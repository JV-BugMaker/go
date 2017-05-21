package main

import (
	"reflect"
)

//value专注于对象实例数据读写
//接口变量会赋值对象 是unaddressable 想要修改目标对象 就必须使用指针

func main() {
	a := 100
	va , vp := reflect.ValueOf(a),reflect.ValueOf(&a).Elem()


	fmt.Println(va.CanAddr(),va.CanSet())
	fmt.Println(vp.CanAddr(),vp.CanSet())
}

