package main
import "fmt"

// Go supports recursive functions.
// Here is the classic factorial example - a function that will call itself until it reaches the base case of 0.

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println("calling recursively!")
	return n * factorial(n - 1)
}

func main() {
	fmt.Println(factorial(7))
}