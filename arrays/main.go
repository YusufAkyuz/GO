package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("--------------------------------------------------")

	var productNames [3]string
	var quantities [3]int
	var prices [3]float32
	var isInStock [3]bool

	productNames[0] = "Jack"
	productNames[1] = "Jill"
	productNames[2] = "John"

	quantities[0] = 24
	quantities[1] = 12
	quantities[2] = 36

	prices[0] = 19.99
	prices[1] = 9.99
	prices[2] = 29.99

	isInStock[0] = true
	isInStock[1] = false
	isInStock[2] = true

	fmt.Println(productNames, reflect.TypeOf(productNames))
	fmt.Println(quantities, reflect.TypeOf(quantities))
	fmt.Println(prices, reflect.TypeOf(prices))
	fmt.Println(isInStock, reflect.TypeOf(isInStock))

	fmt.Println("--------------------------------------------------")

	var productNames2 = [3]string{"Jack", "Jill", "John"}
	var quantities2 = [3]int{24, 12, 36}
	var prices2 = [3]float32{19.99, 9.99, 29.99}
	var isInStock2 = [3]bool{true, false, true}

	fmt.Println(productNames2, reflect.TypeOf(productNames2))
	fmt.Println(quantities2, reflect.TypeOf(quantities2))
	fmt.Println(prices2, reflect.TypeOf(prices2))
	fmt.Println(isInStock2, reflect.TypeOf(isInStock2))

	fmt.Println("--------------------------------------------------")

	productNames4 := [3]string{"Jack", "Jill", "John"}
	quantities4 := [3]int{24, 12, 36}
	prices4 := [3]float32{19.99, 9.99, 29.99}
	isInStock4 := [3]bool{true, false, true}

	fmt.Println(productNames4, reflect.TypeOf(productNames4))
	fmt.Println(quantities4, reflect.TypeOf(quantities4))
	fmt.Println(prices4, reflect.TypeOf(prices4))
	fmt.Println(isInStock4, reflect.TypeOf(isInStock4))

	fmt.Println("--------------------------------------------------")

	var cars = [3]string{"Audi", "Mercedes", "Volkswagen"}
	fmt.Println(cars[0:2], reflect.TypeOf(cars))

	fmt.Println("--------------------------------------------------")

	//==========================================================
	//Slices are dynamic arrays
	//==========================================================
	
	slices := []string{"Audi", "Mercedes", "Volkswagen"}
	fmt.Println(slices, reflect.TypeOf(slices))
	slices = append(slices, "BMW")
	fmt.Println(slices, reflect.TypeOf(slices))

	fmt.Println("--------------------------------------------------")

	var names = []string{"Yusuf", "Ahmet", "Mehmet"}
	fmt.Println(names, reflect.TypeOf(names))
	names = append(names, "Ali")
	fmt.Println(names, reflect.TypeOf(names))

	fmt.Println("--------------------------------------------------")
}
