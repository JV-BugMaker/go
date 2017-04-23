package main

func testPointer(p *int){
	go func() {
		println(p)
	}()

	println(p)
}
func testPointer2(p **int){
	x := 100
	*p = &x
}
func main() {

	x := 100
	p := &x
	testPointer(p)
	testPointer2(&p)
	println(*p)
}


