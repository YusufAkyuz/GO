package main

import (
	"fmt"
)

func main() {
	var age int = 30

	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior.")
	}

	a := 10
	b := 20
	c := 30

	if a > b && a > c {
		fmt.Println("a is the largest number.")
	} else if b > a && b > c {
		fmt.Println("b is the largest number.")
	} else {
		fmt.Println("c is the largest number.")
	}
}
