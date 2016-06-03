package main

import(
	"fmt"
	"log"
	"strings"
	"github.com/VojtechVitek/go-trello"
)

func main() {
	// New Trello Client
	var appKey, token, username string
	fmt.Println("Enter AppKey")
	fmt.Scan(&appKey)
	fmt.Println("Enter Token")
	fmt.Scan(&token)
	trello, err := trello.NewAuthClient(appKey, &token)
	if err != nil {
		log.Fatal(err)
	}

	// User @trello
	fmt.Println("Enter your Trello username")
	fmt.Scan(&username)
	user, err := trello.Member(username)
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
