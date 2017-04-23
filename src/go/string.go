package main

import (
	"unsafe"
	"fmt"
	"reflect"
)

type stringStruct struct {
	//字符串默认是"" 是不可变字节序列，其本身是一个复合结构

	str unsafe.Pointer

	len int
	// ` 定义不做转义处理的原始字符串，支持跨行
}

func toString(bs []byte) string
{
	//以非安全指针类型转换来实现类型变更
	return *(*string)(unsafe.Pointer(&bs))
}
func main() {
	s := "abcdefg"

	s1 := s[:3]
	s2 := s[1:3]
	s3 := s[2:]

	println(s1,s2,s3)

	//notice reflect.StringHeader和string头结构相同
	//unsafe.Pointer用于指针类型转换

	fmt.Printf("%#v\n",(*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Printf("%#v\n",(*reflect.StringHeader)(unsafe.Pointer(&s1)))

	s = "顾嘉伟"

	for i:=0;i<len(s);i++{ //byte
		fmt.Printf("%d:[%c]\n",i,s[i])
	}

	for i,c := range s {
		fmt.Printf("%d:[%c]\n",i,c)
	}
}

