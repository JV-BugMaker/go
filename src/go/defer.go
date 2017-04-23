package main

func main() {
	x ,y := 1,2

	defer func(a int) {
		//1 11 12
		println(a,x,y)
	}(x)

	x +=10
	y += 10
	println(x,y)
}
