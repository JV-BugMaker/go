package main

func main() {
	x := 100

	//同一一个作用域中 最少一个变量重新定义赋值 就会是退化赋值操作
	x,y:=200,"asdasd"
	println(&x)
	//{
	//	//不同作用域就是全新定义赋值
	//	x,y:=200,"abc"
	//	println(x,y)
	//}

	println(x,y)
}


