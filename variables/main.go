package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("--------------------------------------------------")

	var productName string
	var quantity int
	var price float32
	var isInStock bool

	productName = "Jack"
	quantity = 24
	price = 19.99
	isInStock = true

	fmt.Println(productName, reflect.TypeOf(productName))
	fmt.Println(quantity, reflect.TypeOf(quantity))
	fmt.Println(price, reflect.TypeOf(price))
	fmt.Println(isInStock, reflect.TypeOf(isInStock))

	fmt.Println("--------------------------------------------------")

	var productName2, quantity2, price2, isInStock2 = "Jack", 24, 19.99, true

	fmt.Println(productName2, reflect.TypeOf(productName2))
	fmt.Println(quantity2, reflect.TypeOf(quantity2))
	fmt.Println(price2, reflect.TypeOf(price2))
	fmt.Println(isInStock2, reflect.TypeOf(isInStock2))

	fmt.Println("--------------------------------------------------")

	productName4 := "Jack"
	quantity4 := 24
	price4 := 19.99
	isInStock4 := true

	fmt.Println(productName4, reflect.TypeOf(productName4))
	fmt.Println(quantity4, reflect.TypeOf(quantity4))
	fmt.Println(price4, reflect.TypeOf(price4))
	fmt.Println(isInStock4, reflect.TypeOf(isInStock4))

	fmt.Println("--------------------------------------------------")

	fmt.Printf("Product Name: %s\n", productName)
	fmt.Printf("Quantity: %d\n", quantity)
	fmt.Printf("Price: %.2f\n", price)
	fmt.Printf("Is in Stock: %t\n", isInStock)

	fmt.Println("--------------------------------------------------")
}
