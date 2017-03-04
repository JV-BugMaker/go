package main

func test(c byte){
	println(c)
}

func main() {
	//byte alias uint8
	//rune alias int32
	var a byte = 0x11
	var b uint = a
	var c uint8 = a + b

	test(c)

}



