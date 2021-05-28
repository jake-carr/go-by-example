package main

import "fmt"

// Basic examples of Go value types

func main() {
	// Strings which can be added together with +
	fmt.Println("go" + "lang")

	// Integers and floats
	fmt.Println("1+4 =", 1+4)
	fmt.Println("7.4/3.0", 7.4/3.0)

	// Booleans & expected boolean operators
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true, !false)
}