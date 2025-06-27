package main

import "fmt"

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

	fmt.Printf("g type: %T\n", g)                         // Example of using the ID type
	fmt.Printf("g as int: %d\n", int(g))                  // Converting ID to int for demonstration
	fmt.Printf("g as string: %s\n", fmt.Sprintf("%d", g)) // Converting ID to string for demonstration
}
