package main

import (
	"os"
	"strconv"
)

//全局变量定义 不使用并不会抛出错误
var x int
func main() {
	f,err:= os.Open("/tmp/wifi-02-25-2017__17:37:43.log")

	buf := make([]byte,1024)

	//函数返回可能会带错误信息 没有错误的时候 err就是nil
	n,err := f.Read(buf)

	//多变量赋值的时候 首先计算出所有右值 只是赋值
	//赋值操作的时候 必须保证左右类型相同
	x , y :=1,2

	x , y =y+3,x+2

	println(n,err,x,y)

	var c int
	//for条件中并不需要圆括号进行包装起来
	for i :=0;i<10;i++ {
		c++
	}
	//首字母大小写决定其作用域 首字母大写的为导出成员，可被包外引用，而小写则仅能在包内使用
	println(c)

	//和py类似，go也有名为"_"的特殊成员 通常作为忽略占位符使用 可作为表达式左值
	s,_ := strconv.Atoi("12")
	//忽略atoi的err返回值
	println(s)
}
