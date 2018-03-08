// In-place replacement of adjacent duplicates
// Exercise 4.5
package main

import "fmt"

func main() {
	dups := []string{"a", "a", "b", "c", "d", "d", "g", "e", "e", "z"}
	fmt.Println(removeAdjacentDups(dups))
}

func removeAdjacentDups(s []string) []string {
	var prev string
	for i := 0; i < len(s); i++ {
		if s[i] == prev {
			s = append(s[:i-1], s[i+1:]...)
			return removeAdjacentDups(s)
		}

		prev = s[i]
	}

	return s
}
