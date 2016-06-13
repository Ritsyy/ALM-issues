package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/VojtechVitek/go-trello"
	"github.com/google/go-github/github"
	"log"
	"strings"
)

type Configuration struct {
	apiKey   string
	token    string
	userName string
}

type TrelloIssueProvider struct {
	Configuration
	boardId  string
	listName string
}

type GithubIssueProvider struct {
	query string
}

func (t TrelloIssueProvider) FetchData() {
	trello, err := trello.NewAuthClient(t.Configuration.apiKey, &t.Configuration.token)
	if err != nil {
		log.Fatal(err)
	}
	// User @trello
	user, err := trello.Member(t.Configuration.userName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.FullName)

	// @trello Boards
	board, err := trello.Board(t.boardId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("* %v (%v)\n", board.Name, board.ShortUrl)
	// @trello Board Lists
	lists, err := board.Lists()
	if err != nil {
		log.Fatal(err)
	}
	for _, list := range lists {
		if strings.Compare(list.Name, t.listName) == 0 {
			fmt.Println("   - ", list.Name)
			// @trello Board List Cards
			cards, _ := list.Cards()
			for _, card := range cards {
				fmt.Println("      + ", card.Name)
			}
		}
	}
	return
}

func (g GithubIssueProvider) FetchData() {
	client := github.NewClient(nil)
	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	result, _, err := client.Search.Issues(g.query, opts)
	totalCount, _ := json.Marshal(result.Total)
	fmt.Println("Total Count of Issues", string(totalCount))
	issues := result.Issues
	for l, _ := range issues {
		url, _ := json.Marshal(issues[l].URL)
		title, _ := json.Marshal(issues[l].Title)
		fmt.Println("issue url: ", string(url))
		fmt.Println("title: ", string(title))
	}

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}
	return
}

type IssueProvider interface {
	FetchData()
}

func main() {
	var apikey, token, userName, boardId, listName string
	flag.StringVar(&apikey, "apikey", "", "Trello API key")
	flag.StringVar(&token, "token", "", "Trello Token")
	flag.StringVar(&boardId, "boardId", "nlLwlKoz", "Search the board")
	flag.StringVar(&listName, "listName", "Epic Backlog", "Search List from the specific Board")
	flag.StringVar(&userName, "userName", "", "your trello username")
	flag.Parse()
	issueproviders := []IssueProvider{TrelloIssueProvider{Configuration: Configuration{apiKey: apikey, token: token, userName: userName}, boardId: boardId, listName: listName}, GithubIssueProvider{query: "is:open is:issue user:arquillian author:aslakknutsen"}}

	for _, issueprovider := range issueproviders {
		issueprovider.FetchData()
	}
}
