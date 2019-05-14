package twofer

import "fmt"

// Below the result of ShareWith() is passed to standard output
// using fmt.Println, and this is compared against the expected output.
// If they are equal, the test passes.
func ExampleShareWith() {
	h := ShareWith("")
	fmt.Println(h)
	// Output: One for you, one for me.
}
