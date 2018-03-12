// Downloads a poster for a specified movie
// Exercise 4.13
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Movie struct {
	Title  string
	Year   string
	Poster string
}

const omdbURL = "https://omdbapi.com/?t=%s&apikey=a071b57a"

func main() {
	title := strings.Join(os.Args[1:], "+")
	result, err := getMovieData(title)
	if err != nil {
		log.Fatalf("main: %s", err)
	}

	posterData, err := getPoster(result.Poster)
	if err != nil {
		log.Fatalf("main: %s", err)
	}

	writePoster(posterData, result)
}

func getMovieData(title string) (*Movie, error) {
	resp, err := http.Get(fmt.Sprintf(omdbURL, title))
	if err != nil {
		resp.Body.Close()
		return nil, fmt.Errorf("getMovieData(%s) request failed: %s", title, err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getMovieData(%s) bad status: %s", title, err)
	}

	var result Movie
	data, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("getMovieData(%s) unmarshaling failed: %s", title, err)
	}

	return &result, nil
}

func getPoster(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		resp.Body.Close()
		return nil, fmt.Errorf("getPoster(%s) request failed: %s", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getPoster(%s) bad status: %s", url, err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	return data, nil
}

func writePoster(data []byte, movie *Movie) {
	filename := strings.Replace(movie.Title, " ", "_", -1) + ".jpg"
	poster, err := os.Create(filename)
	if err != nil {
		fmt.Errorf("writePoster create failed: %s", err)
	}

	poster.Write(data)
	fmt.Printf("Downloaded poster for %s - %s as %s\n", movie.Title, movie.Year, poster.Name())
}
