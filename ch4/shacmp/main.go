// Returns number of different bits in 2 SHA256 hashes
// Exercise 4.1
package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [8]byte

func main() {

	// Initialize lookup table
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}

	c1 := sha256.Sum256([]byte("Hello, world"))
	c2 := sha256.Sum256([]byte("Hello, sekai"))

	fmt.Printf("SHA 1: %X\n", c1)
	fmt.Printf("SHA 2: %X\n", c2)

	fmt.Printf("Different bit count: %d\n", PopCompare(c1, c2))
}

func PopCompare(x, y [32]byte) int {
	var result int

	// Iterate over bytes
	for i, xval := range x {
		b1 := xval
		b2 := y[i]

		// Compare each bit of byte
		for j := 0; j < 8; j++ {
			if b1&pc[j] != b2&pc[j] {
				result++
			}
		}
	}

	return result
}
