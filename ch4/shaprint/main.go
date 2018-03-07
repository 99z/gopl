// Prints the SHA* of its standard input
// SHA type specified by flags
// Exercise 4.2
package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var sha3 = flag.Bool("384", false, "print SHA384 hash")
var sha5 = flag.Bool("512", false, "print SHA512 hash")

func main() {
	flag.Parse()
	var buffer bytes.Buffer

	for _, val := range os.Args[1:] {
		buffer.WriteString(val)
	}

	if *sha3 {
		sha := sha512.Sum384([]byte(buffer.String()))
		fmt.Printf("SHA384: %X\n", sha)
	} else if *sha5 {
		sha := sha512.Sum512([]byte(buffer.String()))
		fmt.Printf("SHA512: %X\n", sha)
	} else {
		sha := sha256.Sum256([]byte(buffer.String()))
		fmt.Printf("SHA256: %X\n", sha)
	}

}
