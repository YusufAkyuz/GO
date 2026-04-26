package main

import "fmt"

// Structlar, birden fazla alanı tek bir türde gruplayan ve ilişkili verileri bir arada tutan veri yapılarıdır.

type Person struct {
	Name string
	Age  int
}

type Address struct {
	Street string
	City   string
}

type Employee struct {
	Person  // Embedded struct (composition)
	Address // Embedded struct (composition)
	Salary  float64
}

func ToString (person Person) string {
	return fmt.Sprintf("Name: %s, Age: %d", person.Name, person.Age)
}

func main() {
	fmt.Println("--------------------------------------------------")
	p := Person{
		Name: "Ahmet",
		Age:  30,
	}
	fmt.Println(p)

	fmt.Println("--------------------------------------------------")
	var yusuf = Person{Name: "Yusuf", Age: 25}
	fmt.Println(yusuf)

	fmt.Println("--------------------------------------------------")
	var emptyPerson Person
	fmt.Println(emptyPerson)

	fmt.Println("--------------------------------------------------")
	emptyPerson.Name = "Mehmet"
	emptyPerson.Age = 40
	fmt.Println(emptyPerson)

	fmt.Println("--------------------------------------------------")
	var address = Address{
		Street: "123 Main St",
		City:   "Istanbul",
	}
	fmt.Println(address)

	fmt.Println("--------------------------------------------------")
	employee := Employee{
		Person:  Person{Name: "Ayşe", Age: 28},
		Address: Address{Street: "456 Elm St", City: "Ankara"},
		Salary:  5000.00,
	}
	fmt.Println(employee)

	fmt.Println("--------------------------------------------------")
	fmt.Println("Employee ToString:", ToString(p))

	fmt.Println("--------------------------------------------------")
}
