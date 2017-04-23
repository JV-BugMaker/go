package main

func testStruct(){
	type calc struct {
		mult func(x int ,y int) int
	}
	x := calc{
		mult: func(x int, y int) int {
			return x * y
		},
	}

	println(x.mult(2,3))
}
func main() {
	c := make(chan func(int,int) int,2)
	
	c <- func(x,y int) int {
		return x + y
	}

	println((<-c)(1,2))
}