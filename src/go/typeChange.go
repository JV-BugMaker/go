package main

func main() {
	a := 10
	b := byte(a)
	c := a + int(b) //混合类型必须确保类型一致
	println(c)

	//错误 int类型不能作为 bool类型进行true/false判断
	if a {

	}

	//如果转换的目标是指针、单向通道或没有返回值的函数类型，那么必须使用圆括号，以规避造成语法分解错误

	x := 100
	//错误 正确的做法是用括号，让编译器将*int解析为指针类型
	//p := *int(&x)
	//(*int)(p) -->如果没有括号 *(int(p))
	//(<-chan int)(c)
	//(func())(x)
	//(func()int)(x)
	println(p)
}
