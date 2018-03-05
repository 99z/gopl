// Non-recursive version of comma
// Covers exercises 3.10, 3.11
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(comma(strings.Join(os.Args[1:2], "")))
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		buf.WriteString(s)
		return buf.String()
	}

	for i, v := range s {
		if ((len(s) - i) % 3 == 0) && i != 0 {
			buf.WriteByte(',')
		}

		buf.WriteRune(v)
	}

	return buf.String()
}