package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://google.com",
		"https://twitter.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// checkWebsite(link) <- Slow version
		go checkWebsite(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkWebsite(link, c)
		}(l)
	}
}

func checkWebsite(link string, c chan string) {
	_, err := http.Get(link) // Can be improved with go routines

	if err != nil {
		fmt.Println("Website error:", link)
		c <- link
		return
	}

	fmt.Println(link, "website is up!")
	c <- link
}
