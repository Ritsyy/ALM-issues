package main

import(
  "fmt"
  "github.com/google/go-github/github"
)

func main(){
  client := github.NewClient(nil)

  fmt.Println("Open issues authored by aslakuntsen and user is arquillian")

  query := "is:open is:issue user:arquillian author:aslakknutsen"

  opts := &github.SearchOptions{
    ListOptions: github.ListOptions{
      PerPage: 10,
    },
  }

  issues, _, err := client.Search.Issues(query, opts)

  if err != nil {
    fmt.Printf("error: %v\n\n", err)
    } else {
      fmt.Printf("%v\n\n", github.Stringify(issues))
    }

  }
