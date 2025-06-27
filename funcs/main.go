package main

import (
	"errors"
	"fmt"
)

// Function with no parameters and no return value
func sayHello() {
	fmt.Println("Hello!")
}

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function with return value
func add(a int, b int) (int, error) {
	sum := a + b
	if sum < a || sum < b {
		return 0, errors.New("integer overflow")
	}
	return sum, nil
}

// More complex example: returns sum, difference, product, quotient, and error (if division by zero)
func calcAll(a, b int) (sum, diff, prod int, quot float64, err error) {
	sum = a + b
	diff = a - b
	prod = a * b
	if b == 0 {
		err = errors.New("division by zero")
		quot = 0
	} else {
		quot = float64(a) / float64(b)
		err = nil
	}
	return
}

// Function returning multiple values
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	sayHello()
	greet("Alice")
	sum, err := add(3, 4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sum:", sum)
	}
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)

	sum, diff, prod, quot, err := calcAll(5, 10)
	if err != nil {
		fmt.Println(sum)
		fmt.Println(diff)
		fmt.Println(prod)
		fmt.Println(quot)

		fmt.Println("Error:", err)
	} else {
		fmt.Println(sum)
		fmt.Println(diff)
		fmt.Println(prod)
		fmt.Println(quot)
	}

	printSum(1, 2, 3, 4, 5)
	printGreetings("Alice", "Bob", "Charlie")
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func printSum(numbers ...int) {
	total := sum(numbers...)
	fmt.Printf("The sum of %v is %d\n", numbers, total)
}
func printGreeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
func printGreetings(names ...string) {
	for _, name := range names {
		printGreeting(name)
	}
}
