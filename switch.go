package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write ", i, " as ")

	// Basic switch. Note lack of indentation on cases.
	switch i {
	case 1:
		fmt.Println("ONE")
	case 2:
		fmt.Println("TWO")
	case 3:
		fmt.Println("THREE")
	}

	// Use commas to seperate multiple expressions in the same case statement.
	// Providing a default case to a switch is optional.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend!")
	default: // Here using 'default' is more efficient than listing out 5 weekdays.
		fmt.Println("It is not the weekend!")
	}

	// switch without an expression is an alternate way to express if/else logic.
	// Here, we also show how the case expressions can be non-constants:
	t := time.Now()
	switch {
	case t.Hour() < 12: // Non-constant case expression
		fmt.Println("It is before noon!")
	default: // effectively 'Else'
		fmt.Println("It is after noon!")
	}

  // A type switch compares types instead of values. You can use this to discover the type of an interface value. In this example, var t will have the type corresponding to its clause.
	printType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool!!")
		case int:
			fmt.Println("I'm an int!!")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	printType(true)
  printType(1)
  printType("greetings")
}