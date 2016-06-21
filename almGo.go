package main

import (
	"flag"
	"fmt"
	"github.com/VojtechVitek/go-trello"
	"github.com/google/go-github/github"
	"log"
	"strings"
)

type Configuration struct {
	ApiKey   string
	Token    string
	UserName string
}

type TrelloIssueProvider struct {
	Configuration
	BoardId  string
	ListName string
}

type GithubIssueProvider struct {
	Query string
}

type Issue struct {
	title       string
	description string
}

func PrintIssue(issue []Issue) {
	for i := range issue {
		fmt.Println("title: ", string(issue[i].title))
		fmt.Println("issue: ", string(issue[i].description))
		fmt.Println("")
	}
}

func (t TrelloIssueProvider) FetchData(c chan []Issue) {
	var issueArr []Issue
	trello, err := trello.NewAuthClient(t.Configuration.ApiKey, &t.Configuration.Token)
	if err != nil {
		log.Fatal(err)
	}

	// @trello Boards
	board, err := trello.Board(t.BoardId)
	if err != nil {
		log.Fatal(err)
	}

	// @trello Board Lists
	lists, err := board.Lists()
	if err != nil {
		log.Fatal(err)
	}
	for _, list := range lists {
		if strings.Compare(list.Name, t.ListName) == 0 {
			// @trello Board List Cards
			cards, _ := list.Cards()
			for _, card := range cards {
				cardName := card.Name
				description := card.Desc
				issueInstance := Issue{cardName, description}
				issueArr = append(issueArr, issueInstance)
			}
			c <- issueArr
		}
	}
}

func (g GithubIssueProvider) FetchData(c chan []Issue) {
	var issueArr []Issue
	client := github.NewClient(nil)
	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	result, _, err := client.Search.Issues(g.Query, opts)
	issues := result.Issues
	for l, _ := range issues {
		url := issues[l].URL
		title := issues[l].Title
		issueInstance := Issue{*title, *url}
		issueArr = append(issueArr, issueInstance)
	}
	c <- issueArr
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}
}

type IssueProvider interface {
	FetchData(chan []Issue)
}

func main() {
	var tool, apiKey, token, userName, boardId, listName, query string
	flag.StringVar(&tool, "tool", "", "Choose the tool from which you want to search")
	flag.StringVar(&query, "query", "is:open is:issue user:arquillian author:aslakknutsen", "what you want to search on github")
	flag.StringVar(&apiKey, "apiKey", "", "Trello API key")
	flag.StringVar(&token, "token", "", "Trello Token")
	flag.StringVar(&boardId, "boardId", "nlLwlKoz", "Search the board")
	flag.StringVar(&listName, "listName", "Epic Backlog", "Search List from the specific Board")
	flag.StringVar(&userName, "userName", "", "your trello username")
	flag.Parse()
	if tool == "github" {
		var printarr []Issue
		c := make(chan []Issue)
		issueproviders := []IssueProvider{GithubIssueProvider{Query: query}}
		for _, issueprovider := range issueproviders {
			go issueprovider.FetchData(c)
			printarr = <-c
		}
		PrintIssue(printarr)
	} else if tool == "trello" {
		var printarr []Issue
		c := make(chan []Issue)
		issueproviders := []IssueProvider{TrelloIssueProvider{Configuration: Configuration{ApiKey: apiKey, Token: token, UserName: userName}, BoardId: boardId, ListName: listName}}
		for _, issueprovider := range issueproviders {
			go issueprovider.FetchData(c)
			printarr = <-c
		}
		PrintIssue(printarr)
	}
}
