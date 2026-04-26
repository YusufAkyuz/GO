package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("--------------------------------------------------")

	numbers := []int{1, 2, 3, 4, 5}

	for index, number := range numbers {
		fmt.Printf("Index: %d, Number: %d\n", index, number)
	}

	fmt.Println("--------------------------------------------------")

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index: %d, Number: %d\n", i, numbers[i])
	}

	fmt.Println("--------------------------------------------------")

	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum: %d\n", sum)

	fmt.Println("--------------------------------------------------")

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}

	fmt.Println("--------------------------------------------------")

	for i := 0; i < 5; i++ {
		if i == 3 && i == 4 {
			fmt.Println("Breaking the loop at i =", i)
			break
		}
		fmt.Println(i)
	}

	fmt.Println("--------------------------------------------------")

	for i := 0; i < 5; i++ {
		if i%2 == 0 || i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("--------------------------------------------------")

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("--------------------------------------------------")

	index := 0
	names := []string{"Yusuf", "Ahmet", "Mehmet"}

	for index < len(names) {
		fmt.Printf("Index: %d, Name: %s\n", index, names[index])
		index++
	}
}
