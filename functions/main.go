package main

import (
	"fmt"
)

func getFullName(firstName string, lastName string) string {
	return firstName + " " + lastName
}

func add(a int, b int) int {
	return a + b
}

// A function can return multiple values. For example, we can return the full name and the length of the full name.
func getFullNameAndLength(firstName string, lastName string) (string, int) {
	fullName := firstName + " " + lastName
	return fullName, len(fullName)
}

func returnMultiple(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

func sumArray(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func dynamicInput(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}




func main() {
	fmt.Println("--------------------------------------------------")
	fmt.Println(getFullName("Yusuf", "Akyuz"))
	fmt.Println("--------------------------------------------------")
	fullName, length := getFullNameAndLength("Yusuf", "Akyuz")
	fmt.Printf("Full Name: %s, Length: %d\n", fullName, length)
	fmt.Println("--------------------------------------------------")
	fmt.Println("Sum:", add(5, 10))
	fmt.Println("--------------------------------------------------")
	sum, product := returnMultiple(5, 10)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)
	fmt.Println("--------------------------------------------------")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum of array: %d\n", sumArray(numbers))
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Sum of dynamic input: %d\n", dynamicInput(1, 2, 3, 4, 5, 6))
	fmt.Printf("Sum of dynamic input: %d\n", dynamicInput(1, 2, 3, 4, 5, 6, 7, 8, 9, 0))
	fmt.Println("--------------------------------------------------")
}
