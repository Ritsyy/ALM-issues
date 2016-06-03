package main

import(
	"fmt"
	"log"
	"github.com/VojtechVitek/go-trello"
	"strings"
)

func main() {
	// New Trello Client
	appKey := "//your appkey"
	token := "your token"
	trello, err := trello.NewAuthClient(appKey, &token)
	if err != nil {
		log.Fatal(err)
	}

	// User @trello
	user, err := trello.Member("your TrelloUserName")
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
			if strings.Compare(board.Name, "AtomicOpenShift Roadmap")==0{
				fmt.Printf("* %v (%v)\n", board.Name, board.ShortUrl)
				// @trello Board Lists
				lists, err := board.Lists()
				if err != nil {
					log.Fatal(err)

				}
				for _, list := range lists {
					if strings.Compare(list.Name, "Epic Backlog")==0{
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
