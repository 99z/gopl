// Reports the frequency of each word in an input txt file
// Exercise 4.9
package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

func main() {
	file, err := os.Open("/lib/gettysburg")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	
	defer file.Close()

	in := bufio.NewScanner(file)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		w := in.Text()

		counts[w]++
	}

	fmt.Printf("word\tcount\n")
	for c, n := range counts {
		fmt.Printf("%s\t%d\n", c, n)
	}
}
