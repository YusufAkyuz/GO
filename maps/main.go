package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("--------------------------------------------------")

	// Key-value pairs. Keys are unique, values can be duplicated.
	// FOr keys equals the string type and value is int type we use map[string]int
	var names map[string]int
	names = make(map[string]int)

	names["Yusuf"] = 30
	names["Ahmet"] = 25
	names["Mehmet"] = 35

	fmt.Println(names, reflect.TypeOf(names))
	fmt.Println("Yusuf key's value is: ", names["Yusuf"])
	fmt.Println("Ahmet key's value is: ", names["Ahmet"])
	fmt.Println("Mehmet key's value is: ", names["Mehmet"])

	fmt.Println("--------------------------------------------------")

	// We can also declare and initialize a map in one line.
	names2 := map[string]int{
		"Yusuf":  30,
		"Ahmet":  25,
		"Mehmet": 35,
	}

	fmt.Println(names2, reflect.TypeOf(names2))
	fmt.Println("Yusuf key's value is: ", names2["Yusuf"])
	fmt.Println("Ahmet key's value is: ", names2["Ahmet"])
	fmt.Println("Mehmet key's value is: ", names2["Mehmet"])

	fmt.Println("--------------------------------------------------")

	// We can also declare and initialize a map using make function.
	names3 := make(map[string]int)
	names3["Yusuf"] = 30
	names3["Ahmet"] = 25
	names3["Mehmet"] = 35

	fmt.Println(names3, reflect.TypeOf(names3))
	fmt.Println("Yusuf key's value is: ", names3["Yusuf"])
	fmt.Println("Ahmet key's value is: ", names3["Ahmet"])
	fmt.Println("Mehmet key's value is: ", names3["Mehmet"])

	fmt.Println("--------------------------------------------------")

	//Delete a key-value pair from the map
	delete(names3, "Ahmet")
	fmt.Println(names3, reflect.TypeOf(names3))
	fmt.Println("Yusuf key's value is: ", names3["Yusuf"])
	fmt.Println("Ahmet key's value is: ", names3["Ahmet"])

}
