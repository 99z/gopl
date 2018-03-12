// Modified version of issues that includes age categories
// Exercise 4.10
package main

import (
	"./github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	var age string
	for _, item := range result.Items {
		createdHours := time.Since(item.CreatedAt).Hours()
		switch {
		case createdHours < 730:
			age = "<1 month"
		case createdHours < 8760:
			age = "<1 year"
		case createdHours > 8760:
			age = ">1 year"
		}
			fmt.Printf("#%-5d %s %9.9s %.55s\n",
				item.Number, age, item.User.Login, item.Title)
		}
}
