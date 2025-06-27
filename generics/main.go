package main

import (
	"fmt"
)

// Define a custom type
type MyNumber int

// Number interface using ~ for underlying types
type Number interface {
	~int | ~float64
}

// Generic function using type parameter T
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func Sum[T Number](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

// Function using empty interface (void interface)
func PrintAnySlice(s []interface{}) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	ints := []int{1, 2, 3}
	strs := []string{"a", "b", "c"}
	myNums := []MyNumber{10, 20, 30}

	fmt.Println("Using generics:")
	PrintSlice(ints)
	PrintSlice(strs)
	PrintSlice(myNums)

	fmt.Println("\nUsing void interfaces:")
	anys := []interface{}{1, "hello", 3.14, true}
	PrintAnySlice(anys)

	PrintSlice(anys)

	fmt.Println("\nUsing generic sum function:")
	sumInts := Sum(ints)
	sumFloats := Sum([]float64{1.1, 2.2, 3.3})
	sumMyNums := Sum(myNums)
	fmt.Println("Sum of ints:", sumInts)
	fmt.Println("Sum of floats:", sumFloats)
	fmt.Println("Sum of MyNumbers:", sumMyNums)
	fmt.Println("\nUsing generic sum function with type parameter:")
}
