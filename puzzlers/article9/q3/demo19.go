package main

import "fmt"

func main() {
	var m map[string]int

	key := "two"
	elem, ok := m["two"]
	if !ok {
		fmt.Printf("ok is %t\n", ok)
	}
	fmt.Printf("%q,%d\n", key, elem)
	fmt.Printf("The length of nil map: %d\n", len(m))

	fmt.Printf("Delete the key-element pair by key %q...\n", key)
	delete(m, key)

	elem = 2
	fmt.Printf("Add a key-element pair to a nil map...")

	m["two"] = elem
}
