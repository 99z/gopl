// Version of rotate that operates on a single pass
// Exercise 4.4
package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(a[:], 36 % len(a)))
	fmt.Println(a)
}

func rotate(s []int, amt int) []int{
	rotateMembers := s[:amt]
	rotateBase := s[amt:]
	return append(rotateBase, rotateMembers...)
}