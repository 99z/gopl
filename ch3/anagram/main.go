// Anagram tester
// Exercise 3.12
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("usage: main string1 string2")
		return
	}

	fmt.Println(anagram(strings.Join(os.Args[1:2], ""), strings.Join(os.Args[2:3], "")))
}

func anagram(s1, s2 string) bool {
	if (len(s1) == 0) || (len(s2) == 0) {
		fmt.Println("Missing argument")
		os.Exit(1)
	} else if len(s1) != len(s2) {
		fmt.Println("Strings are of different lengths!")
		os.Exit(1)
	} else if s1 == s2 {
		return true
	}

	var count int

	for _, v := range s1 {
		for _, w := range s2 {
			if v == w {
				count++
				break
			}
		}
	}

	if count == len(s1) {
		return true
	} else {
		return false
	}

}
