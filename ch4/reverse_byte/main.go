// In-place reverse of []byte slice
// Exercise 4.7
package main

import "fmt"

func main() {
	a := []byte("Hello, world")
	fmt.Println(string(reverse(a)))
}

func reverse(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}
