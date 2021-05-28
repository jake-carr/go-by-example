package main
import "fmt"

// A function that takes two ints and returns their sum as an int:
func plus(a int, b int) int {
	// Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
	return a + b
}

// When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func greet(name string) string {
	return "Hello " + name + " from this Nifty Go Program"
}

// Call a function just as you’d expect, with name(args).
func main() {
	result := plus(1393, 3901)
	fmt.Println("1393 + 3901 =", result)

	result = plusPlus(122, 4891, 53834)
	fmt.Println("122 + 4891 + 53834 =", result)

	greeting := greet("jake")
	fmt.Println(greeting)
}
