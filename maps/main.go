package main

import (
	"fmt"
)

func main() {
	// Creating a map with string keys and int values
	ages := make(map[string]int)

	// Adding elements
	ages["Alice"] = 30
	ages["Bob"] = 25

	// Accessing elements
	fmt.Println("Alice's age:", ages["Alice"])

	// Checking if a key exists
	age, exists := ages["Charlie"]
	if exists {
		fmt.Println("Charlie's age:", age)
	} else {
		fmt.Println("Charlie not found in map")
	}

	// Iterating over a map
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Deleting an element
	delete(ages, "Bob")
	fmt.Println("After deleting Bob:", ages)

	// Example: Using a FOR loop to print all names in the map
	names := []string{}
	for name := range ages {
		names = append(names, name)
	}
	for i := 0; i < len(names); i++ {
		fmt.Println("Person:", names[i])
	}

	// Example: Getting the length of the map
	fmt.Println("Number of people in map:", len(ages))
}
