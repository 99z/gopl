// Modified charcount to count letters, numbers
// Exercise 4.8
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	types := make(map[string]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		
		if err == io.EOF {
			break
		}
		
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		
		if unicode.IsLetter(r) {
			types["letter"]++
		} else if unicode.IsNumber(r) {
			types["number"]++
		}
		
		counts[r]++
		utflen[n]++
	}
	
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	
	fmt.Printf("\ntype\tcount\n")
	for t, n := range types {
		fmt.Printf("%s\t%d\n", t, n)
	}
	
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}