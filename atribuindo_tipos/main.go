package main

type ID int

var (
	b bool    = true
	c int     = 10
	d float64 = 3.14
	f string  = "Goodbye, World!"
	g ID      = 1
)

func main() {
	a := "X"

	println(a)
	if b {
		println("b is true")
	} else {
		println("b is false")
	}
	println("c:", c)
	println("d:", d)
	println("f:", f)
	println("g:", g)
}
