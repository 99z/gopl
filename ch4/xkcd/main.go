// Populates a local database of xkcd comic information
// Allows for searching the local DB for keywords in the comics
// Exercise 4.12
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const xkcdURL = "https://xkcd.com/%d/info.0.json"
const maxID = 1966

var dbFile = flag.Bool("db", false, "path to marshaled comic db file")

type Comic struct {
	Title      string
	Transcript string
	Img        string
	Num        int
}

func main() {
	flag.Parse()
	if *dbFile {
		comicDB, err := populateDB(strings.Join(os.Args[2:3], ""))
		if err != nil {
			log.Fatalf("main: %s", err)
		}

		results := search(os.Args[3:], comicDB)
		printResults(results)
	} else {
		comicDB, err := downloadRange()
		if err != nil {
			log.Fatalf("main: %s", err)
		}

		writeDB(comicDB)
		fmt.Println("Populated xkcd.json, now run with -db flag and search terms")
	}
}

func printResults(results []*Comic) {
	for _, comic := range results {
		fmt.Printf("https://xkcd.com/%d/\n%s\n", comic.Num, comic.Transcript)
	}
}

func downloadRange() (map[int]*Comic, error) {
	comicDB := make(map[int]*Comic)
	for i := 1; i <= maxID; i++ {
		if i == 404 {
			continue
		}

		comic, err := downloadComic(i)
		if err != nil {
			return nil, fmt.Errorf("downloadRange(%d) failed: %s", i, err)
		}

		comicDB[i] = comic
	}

	return comicDB, nil
}

func downloadComic(id int) (*Comic, error) {
	resp, err := http.Get(fmt.Sprintf(xkcdURL, id))
	if err != nil {
		resp.Body.Close()
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("downloadComic(%d) failed: %s", id, resp.Status)
	}

	var result Comic
	data, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("JSON unmarshaling failed: %s", err)
	}

	return &result, nil
}

func writeDB(comicDB map[int]*Comic) (string, error) {
	data, err := json.MarshalIndent(comicDB, "", "	")
	if err != nil {
		return "", fmt.Errorf("writeDB: marshal failed: %s", err)
	}

	file, writeErr := os.Create("xkcd.json")
	if writeErr != nil {
		return "", fmt.Errorf("writeDB: write failed: %s", writeErr)
	}
	defer file.Close()

	_, copyErr := io.Copy(file, bytes.NewReader(data))
	if copyErr != nil {
		return "", fmt.Errorf("writeDB: copy failed: %s", copyErr)
	}

	return "xkcd.json", nil
}

func populateDB(path string) (map[int]*Comic, error) {
	var comics map[int]*Comic
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("populateDB(%s) failed: %s", path, err)
	}

	json.Unmarshal(data, &comics)
	return comics, nil
}

func search(terms []string, comicDB map[int]*Comic) []*Comic {
	var results []*Comic
	for _, term := range terms {
		for i := 1; i < maxID; i++ {
			if i == 404 {
				continue
			}

			if strings.Contains(comicDB[i].Title, term) ||
				strings.Contains(comicDB[i].Transcript, term) {
				results = append(results, comicDB[i])
			}
		}
	}

	return results
}
