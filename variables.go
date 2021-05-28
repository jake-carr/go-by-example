package main

import "fmt"

// In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.

func main() {
	// var declares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	// can declare multiple at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables
	var d = true
	fmt.Println(d)

	// Vars declared without a coreesponding initialization are zero-valued. For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)

	// := syntax is shorthand for declaring & initalizing a variable; in this case, var f string = "apple":
	f := "apple"
	fmt.Println(f)
}