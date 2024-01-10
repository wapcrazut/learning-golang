package main

import "fmt"

// Version without interfaces
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetingSpanish(eb)
	printGreetingEnglish(sb)
}

func (englishBot) getGreeting() string {
	// Please imagine a very specific logic here
	return "Hello"
}

func (spanishBot) getGreeting() string {
	// Please imagine a very specific logic here
	return "Hola"
}

func printGreetingSpanish(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreetingEnglish(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
