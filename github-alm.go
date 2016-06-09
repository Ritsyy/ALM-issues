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

	result, _, err := client.Search.Issues(query, opts)
	totalCount, _ := json.Marshal(result.Total)
	fmt.Println(string(totalCount))
       issues := result.Issues
	for l, _ := range issues {
		url, _ := json.Marshal(issues[l].URL)
		title, _ := json.Marshal(issues[l].Title)
               labels, _ := json.Marshal(issues[l].Labels[0].Name)
		fmt.Println("issue url: ", string(url))
		fmt.Println("title: ", string(title))
               fmt.Println("labels: ", string(labels))
	}

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}
}
