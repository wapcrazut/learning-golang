package main

import "fmt"

// Declare interface with method
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// Please imagine a very specific logic here
	return "Hello"
}

func (spanishBot) getGreeting() string {
	// Please imagine a very specific logic here
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
