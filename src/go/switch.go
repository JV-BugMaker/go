package main

func main() {
	switch x:=4;x{
	case 4:
		x+=2
		print(x)
		//不需要进行条件匹配 按照源码顺序进行到下一个case中执行print
		//但是可以使用break进行阻止
		fallthrough
	case 5:
		x+=3
	default:
		print(x+2)
	}
	print(x) //x是局部变量
}