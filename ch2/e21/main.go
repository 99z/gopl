package main

import (
	"fmt"
	"./tempconv"
)

func main() {
	fmt.Println(tempconv.CtoK(tempconv.BoilingC))
	fmt.Println(tempconv.KtoC(tempconv.CtoK(tempconv.FreezingC)))
}
