package main

import (
	"reflect"
	"fmt"
)
type X struct {}

func (X) Test(x,y int) (int,error){
	return x + y ,fmt.Errorf("err:%d",x+y)
}

//变参
func (X) Format(s string,a ...interface{}) string{
	return fmt.Sprintf(s,a...)
}

func main() {
	var a X
	v := reflect.ValueOf(&a)
	m := v.MethodByName("Test")

	//参数配置
	in := []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	}
	//使用CallSlice方便一些
	out := m.Call(in) //调用函数 传递参数
	for _,v:=range out {
		fmt.Println(v)
	}
}
