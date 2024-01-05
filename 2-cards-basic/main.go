package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard()}
	cards = append(cards, "Seven of Diamonds")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
