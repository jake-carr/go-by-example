package main

import "fmt"

// Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

func main() {
	// Unlike arrays, slices are typed only by the elements they contain, not the number of elements.
	// Use the builtin method make to create an empty slice.

	demoSlice := make([]string, 3) // Makes a slice of strings of length 3, initially zero-valued.
	fmt.Println("zero-valued slice of strings:", demoSlice)

	// Slices can set and get just like arrays. Builtin len also works on slices.
	demoSlice[0] = "a"
	demoSlice[1] = "b"
	demoSlice[2] = "c"
	fmt.Println("set:", demoSlice)
	fmt.Println("get:", demoSlice[2])
	fmt.Println("len:", len(demoSlice))

	/*
	In addition to these basic operations, slices support several more that make them richer than arrays.
	One is the builtin append, which returns a slice containing one or more new values.
	Note that we need to accept a return value from append as we may get a new slice value.
	*/

	demoSlice = append(demoSlice, "d")
	demoSlice = append(demoSlice, "e", "f")
	fmt.Println("appended: ", demoSlice)

	// Slices can also be copy’d. Here we create an empty slice c of the same length as our original slice and copy into c from original.
	c := make([]string, len(demoSlice))
	copy(c, demoSlice) // copy(to, from)
	fmt.Println("copy: ", c)

	// Slices also support a "slice" operator with the syntax sliceName[low:high]. For example, this gets a slice of the elements demoSlice[2], demoSlice[3] and demoSlice[4]
	partial := demoSlice[2:5]
	fmt.Println("subslice 1:", partial)

	// Slice up to, but exclude, index 5
	partial = demoSlice[:5]
	fmt.Println("subslice 2:", partial)

	// Slice up from and including index 2
	partial = demoSlice[2:]
	fmt.Println("subslice 3:", partial)

	// We can declare and initalize a variable for slice in a single line as well.
	anotherSlice := []string{"g", "h", "i"}
  fmt.Println("declared:", anotherSlice)

	// Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
			innerLen := i + 1
			twoD[i] = make([]int, innerLen)
			for j := 0; j < innerLen; j++ {
					twoD[i][j] = i + j
			}
	}
	fmt.Println("2d: ", twoD)

	// Note that while slices are different TYPES than arrays, they are rendered similarly by fmt.Println.
}