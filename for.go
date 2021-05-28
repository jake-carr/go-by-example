package main

import "fmt"

// for is Go's only looping construct

func main() {
	i := 1 // shorthand for: var i int = 1

	// Basic for loop with single condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Initial/condition/after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without a condition will loop repeatedly, until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("looping")
		break
	}

	// you can also continue to the next iteration of the loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
				continue
		}
		fmt.Println(n)
	}
}