package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
)

func main() {
	client := github.NewClient(nil)

	fmt.Println("Open issues authored by aslakuntsen and user is arquillian")

	query := "is:open is:issue user:arquillian author:aslakknutsen"

	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{
			PerPage: 5,
		},
	}

	issues, _, err := client.Search.Issues(query, opts)
	res, _ := json.Marshal(issues.Total)
	fmt.Println(string(res))
	for l, _ := range issues.Issues {
		res1, _ := json.Marshal(issues.Issues[l].URL)
		res2, _ := json.Marshal(issues.Issues[l].Title)
                res3, _ := json.Marshal(issues.Issues[l].Labels[0].Name)
		fmt.Println("issue url: ", string(res1))
		fmt.Println("title: ", string(res2))
                fmt.Println("labels: ", string(res3))
	}

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}
}
