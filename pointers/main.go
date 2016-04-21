package main

import "fmt"

func main() {
	// without pointer
	var x int
	x = 10
	fmt.Println(x)
	fmt.Println(&x)

	// with pointer
	var num *int
	val := new(int)

	num = new(int)
	*num = x

	val = &x

	fmt.Println("=== pointer num ===")
	fmt.Println(*num) // print a value of x
	fmt.Println(num)  // print address of x

	fmt.Println("=== pointer val ===")
	fmt.Println(*val) // print a value of x
	fmt.Println(val)  // print address of x
}
