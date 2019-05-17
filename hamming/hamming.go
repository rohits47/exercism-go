// Package hamming contains methods to assist with calculating the Hamming
// distance between different objects.
package hamming

import "errors"

// Distance calculates the Hamming distance between two DNA strands,
// given as strings.
func Distance(a, b string) (int, error) {
	// return 0, nil
	return -1, errors.New("can't calculate Hamming distance")
}
