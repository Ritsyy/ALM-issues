package main

import(
	"fmt"
	"log"
	"strings"
	"flag"
	"github.com/VojtechVitek/go-trello"
)

func main() {
	// New Trello Client
	var apikey, token, userName, boardName, listName string
	flag.StringVar(&apikey, "apikey","", "Trello API key")
	flag.StringVar(&token, "token","", "Trello Token")
	flag.StringVar(&boardName, "boardName", "AtomicOpenShift Roadmap", "Search the board")
	flag.StringVar(&listName, "listName"	, "Epic Backlog","Search List from the specific Board")
	flag.StringVar(&userName, "userName","", "your trello username")
	flag.Parse()
	trello, err := trello.NewAuthClient(apikey, &token)
	if err != nil {
		log.Fatal(err)
	}

	// User @trello
	user, err := trello.Member(userName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.FullName)

	// @trello Boards
	boards, err := user.Boards()
	if err != nil {
		log.Fatal(err)
	}
	if len(boards) > 0 {
		for i:=0; i<len(boards);i++{
			board := boards[i]
			if strings.Compare(board.Name, boardName)==0{
				fmt.Printf("* %v (%v)\n", board.Name, board.ShortUrl)
				// @trello Board Lists
				lists, err := board.Lists()
				if err != nil {
					log.Fatal(err)
				}
				for _, list := range lists {
					if strings.Compare(list.Name, listName)==0{
						fmt.Println("   - ", list.Name)
						// @trello Board List Cards
						cards, _ := list.Cards()
						for _, card := range cards {
							fmt.Println("      + ", card.Name)
						}
					}
				}
			}
		}
	}
}
