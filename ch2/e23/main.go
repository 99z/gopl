package main

import (
	"fmt"
	"os"
	"strconv"
	"./popcount"
)

func main() {
	for _, arg := range os.Args[1:] {
		u, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			os.Stderr.WriteString("Error parsing arguments")
			os.Exit(1)
		}

		fmt.Println(popcount.PopCount(u))
	}
}
