package main

import (
	"fmt"
	"os"
	"strconv"
	"./weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		w, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		p := weightconv.Pounds(w)
		k := weightconv.Kilograms(w)
		fmt.Printf("%s = %s, %s = %s\n",
			p, weightconv.PToK(p), k, weightconv.KToP(k))
	}
}