// Package twofer provides functions around the concept of "two for one".
package twofer

import "fmt"

// ShareWith returns a string about sharing, customized for the given
// name string.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
