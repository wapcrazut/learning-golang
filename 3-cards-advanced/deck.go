package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of deck
// Which will be a slice of string
type deck []string

func newDeckFromFile(filename string) deck {

	bs, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}

func newDeck() deck {
	cards := deck{}

	cardTypes := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, t := range cardTypes {
		for _, v := range cardValues {
			cards = append(cards, v+" of "+t)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {
	// Old versions of go needed to seed the number
	//source := rand.NewSource(time.Now().UnixNano())
	//r := rand.New(source)

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func deal(d deck, handSize int) (deck, deck) {

	println("Dealing cards!")
	cardsToDeal := d[:handSize]
	restOfDeck := d[handSize:]

	return cardsToDeal, restOfDeck
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
