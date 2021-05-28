package main

import "fmt"

// In Go, an array is a numbered sequence of elements of a specific length.

func main() {
	// Create an array A that will hold exactly 5 ints; the type of elements and length are both part of an array's type. By default, an aray is zero-valued, which of ints, means 0s.
	var a [5]int
	fmt.Println("empty: ", a);

	// We can set a value at an index using the array[index] = value syntax, and get a value with array[index], same as in JavaScript.
	a[4] = 100
	fmt.Println("set: ", a);
	fmt.Println("get", a[4]);

	// The builtin function "len" returns the length of an array
	fmt.Println("array length:", len(a))

	// Syntax to declare and initalize an array in one line:
	b := [5]int{1, 2, 3, 4, 5} // var b is always an [array of five ints], and it begins with the values {1-5}
	fmt.Println("declared:", b)

	// Array types are one-dimensional but you can compose types to build multi-dimensional data structures.
	var twoDimensional [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimensional[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoDimensional)

}