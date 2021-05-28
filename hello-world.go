package main

/*
Go provides various in-built packages to users so they can ease into coding with pre-defined basic functionality packages. One of such packages in the “fmt” package. fmt stands for the Format package. This package allows to format basic strings, values, or anything and print them or collect user input from the console, or write into a file using a writer or even print customized fancy error messages. This package is all about formatting input and output.
*/

import "fmt"

func main() {
	fmt.Println("hello world")
}

// $ go run hello-world.go