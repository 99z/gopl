// Fetch prints the content found at each specified URL
// Uses io.Copy to write content to Stdout
// Prepends 'http://' if not provided
// Includes status code
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http:////") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		
		io.Copy(os.Stdout, resp.Body)
		fmt.Printf("Server responded with:\t%s\n", resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
		resp.Body.Close()
	}
}
