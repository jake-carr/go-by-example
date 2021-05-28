package main

import "fmt"

// Maps are Gos built-in associative data types, sometimes called hashes or dicts in other languages.

func main() {
	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	m["key1"] = 5;
	m["key2"] = 11;

	// Printing a map will show all of its key/value pairs.
	fmt.Println("map: ", m)

	// Get a value for a key with name[key]. The builtin len returns the number of key/value pairs when called on a map.
	val1 := m["key1"]
	fmt.Println("value one:", val1)
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map.
	delete(m, "key2");
	fmt.Println("map after deleting a key: ", m)

	/*
	The optional second return value when getting a value from a map indicates if the key was present in the map.
	This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	*/

	_, present := m["key2"]
	fmt.Println("is key 2 present?", present)

	// You can also declare and initialize a new map in the same line with this syntax:
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}