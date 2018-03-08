// In-place squashing of multiple UTF8 spaces to one ASCII
// Exercise 4.6
package main

import (
	"fmt"
	"unicode"
)

func main() {
	spaces := []byte("Hello, \n\t wor\n\tld")
	spacesSingle := []byte("Hello, \n\t wor\n\tld")
	fmt.Println(string(squashSpaces(spaces)))
	fmt.Println(string(squashSpacesSinglePass(spacesSingle)))
}

// Recursive solution, not O(n), not single-pass
func squashSpaces(b []byte) []byte {
	var count, begin, end int
	for i := 0; i < len(b); i++ {
		if unicode.IsSpace(rune(b[i])) {
			count++
			if count == 1 {
				begin = i
			}

			end = i
		} else if count > 1 {
			b = append(b[:begin], b[end+1:]...)
			b = append(b, 0)
			copy(b[begin+1:], b[begin:])
			b[begin] = 32
			return squashSpaces(b)
		} else if count == 1 {
			count, begin, end = 0, 0, 0
		}
	}

	return b
}

// Non-recursive, O(n) solution, single-pass
func squashSpacesSinglePass(b []byte) []byte {
	var insertionPtr, charPtr, count, removedChars int
	for _, val := range b {
		if !unicode.IsSpace(rune(val)) {
			b[insertionPtr] = val
			insertionPtr++
			charPtr++
			count = 0
		} else {
			if count == 0 {
				b[insertionPtr] = val
				insertionPtr++
			} else if count > 0 {
				if count == 1 {
					b[insertionPtr-1] = 32
				}
				removedChars++
			}
			count++
			charPtr++
		}
	}

	return b[:len(b)-removedChars]
}
