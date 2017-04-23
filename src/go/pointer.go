package main

func main() {
	 x := 10

	var p *int = &x

	*p += 20

	print(p,"|",*p)

	a := struct {
		c int
	}{}

	a.c = 10

	pp := &a
	//
	pp.c += 100

	println("|",pp.c)
}