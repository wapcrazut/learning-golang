package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type customWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Example on how to use byte slices
	//buffer := make([]byte, 99999)
	//resp.Body.Read(buffer)
	//fmt.Println(string(buffer))

	// Using the Write interface
	//io.Copy(os.Stdout, resp.Body)

	// Using custom Write implementation
	cw := customWriter{}
	io.Copy(cw, resp.Body)
}

// Implements default Write function, meaning it becomes part of the types that implements the method.
func (customWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Bytes written:", len(bs))
	return len(bs), nil
}
