package main

import "fmt"

func main() {
	// Create a slice of integers
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("Original slice:", nums)

	// Append elements
	nums = append(nums, 60, 70)
	fmt.Println("After append:", nums)

	// Slicing
	sub := nums[1:4]
	fmt.Println("Sliced (1:4):", sub)
	fmt.Println("Sliced (2:end):", nums[2:])
	fmt.Println("Sliced (start:2):", nums[:2])
	fmt.Println("Sliced (1:3):", nums[1:3])
	fmt.Println("Sliced (2:4):", nums[2:4])
	fmt.Println("Sliced (1:):", nums[1:])
	fmt.Println("Sliced (:1):", nums[:1])
	fmt.Println("Sliced (:3):", nums[:3])
	fmt.Println("Sliced (3:):", nums[3:])

	// Length and capacity
	fmt.Println("Length:", len(nums), "Capacity:", cap(nums))

	// Make a slice with make()
	// Create a slice of strings with a predefined length of 3
	letters := make([]string, 3)
	letters[0] = "a"
	letters[1] = "b"
	letters[2] = "c"
	fmt.Println("Letters:", letters)

	// Copy a slice
	copyLetters := make([]string, len(letters))
	copy(copyLetters, letters)
	fmt.Println("Copied letters:", copyLetters)

	// Example: Capacity and append
	// Preallocate a slice with capacity for efficiency if you know the size in advance
	prealloc := make([]int, 0, 10) // length 0, capacity 10
	fmt.Println("Preallocated slice:", prealloc, "Length:", len(prealloc), "Capacity:", cap(prealloc))

	// Append elements up to capacity
	for i := 1; i <= 10; i++ {
		prealloc = append(prealloc, i)
	}
	fmt.Println("After appending 10 elements:", prealloc, "Length:", len(prealloc), "Capacity:", cap(prealloc))

	// If you append beyond the capacity, Go automatically allocates a new underlying array with larger capacity
	prealloc = append(prealloc, 11)
	fmt.Println("After appending beyond capacity:", prealloc, "Length:", len(prealloc), "Capacity:", cap(prealloc))

	// Best practice: If you know the number of elements, use make with capacity to avoid repeated allocations
	// If you don't know, just use append and let Go handle resizing

	// Example: Slicing does not copy data, it creates a new slice header referencing the same underlying array
	original := []int{1, 2, 3, 4, 5}
	sliceA := original[1:4] // [2 3 4]
	sliceA[0] = 99          // Modifies original[1]
	fmt.Println("Original after modifying sliceA:", original)
	fmt.Println("sliceA:", sliceA)

	// If you want a true copy, use copy()
	copied := make([]int, len(original))
	copy(copied, original)
	copied[0] = 100
	fmt.Println("Original after modifying copied:", original)
	fmt.Println("Copied:", copied)
}
