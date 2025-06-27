package main

import "fmt"

func main() {

	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 15
	meuArray[2] = 32

	println(meuArray[0], meuArray[1], meuArray[2])
	println("Array length:", len(meuArray))
	println("Array capacity:", cap(meuArray))
	println("Array type:", fmt.Sprintf("%T", meuArray))                 // Print the type of the array
	println("Array as string:", fmt.Sprintf("%v", meuArray))            // Print the array as a string
	println("Array as formatted string:", fmt.Sprintf("%+v", meuArray)) // Print the array with field names

	for i := 0; i < len(meuArray); i++ {
		println("meuArray[", i, "]: ", meuArray[i])
	}

	for i, v := range meuArray {
		println("meuArray[", i, "]: ", v)
	}
}
