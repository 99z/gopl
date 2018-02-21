// Fetchall fetches URLs in parallel and reports their times and sizes
// Results can be viewed in the contents of fetchall.out
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	f, err := os.Create("fetchall.out")
	if err != nil {
		fmt.Sprint(err)
		os.Exit(1)
	}
	
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	
	for range os.Args[1:] {
		f.WriteString(<-ch)
	}
	
	f.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel named ch
		return
	}
	
	// Just want byte count, so don't do anything with the body
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}
