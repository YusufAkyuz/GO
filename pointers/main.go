package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var y *int = &x

	fmt.Println("--------------------------------------------------")

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of y (address of x):", y)
	fmt.Println("Value pointed to by y (value of x):", *y)

	*y = 20
	fmt.Println("Updated value of x after changing through pointer y:", x)

	fmt.Println("--------------------------------------------------")

	var a int = 5
	var b int
	b = a
	var p *int = &a
	
	fmt.Println("Value of a:", a)
	fmt.Println("Value of b (copy of a):", b)
	fmt.Println("Address of a:", &a)
	fmt.Println("Address of b:", &b)
	fmt.Println("Value of p (address of a):", p)
	fmt.Println("Value pointed to by p (value of a):", *p)

	fmt.Println("--------------------------------------------------")

	*p = 777
	fmt.Println("Value of a:", a)
	fmt.Println("Value of b (copy of a):", b)
	fmt.Println("Value pointed to by p (value of a):", *p)

	fmt.Println("--------------------------------------------------")
	
	var num int = 10
	fmt.Println("Value of a before increment function:", num)
	increment(num)
	fmt.Println("Value of a after increment function (should be unchanged):", num)

	fmt.Println("--------------------------------------------------")
	fmt.Println("Value of a before incrementPointer function:", num)
	incrementPointer(&num)
	fmt.Println("Value of a after incrementPointer function (should be incremented):", num)

	fmt.Println("--------------------------------------------------")

	//Array ve slicelar doğası gereği referans tipler olduğu için func'daki değişiklikler
	//main func'daki slice'ı da değiştirecektir.
	var numbers = []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", numbers)

	modifySlice(numbers)
	fmt.Println("Modified slice after function call (should be changed):", numbers)
	fmt.Println("After modifySlice:", numbers)
}

func increment(x int) {
	x = x + 10
}

func incrementPointer(x *int) {
	*x = *x + 10
}

func modifySlice(s []int) {
	for i := range s {
		s[i] = s[i] * 2
	}
}
