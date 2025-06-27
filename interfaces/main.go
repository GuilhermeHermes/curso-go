package main

import "fmt"

// Define an interface
type Speaker interface {
	Speak() string
}

// Implement the interface with a struct
type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

// Another struct implementing the interface
type Dog struct {
	Breed string
}

func (d Dog) Speak() string {
	return "Woof! I am a " + d.Breed
}

func main() {
	var s Speaker

	s = Person{Name: "Alice"}
	fmt.Println(s.Speak())

	s = Dog{Breed: "Labrador"}
	fmt.Println(s.Speak())
}
