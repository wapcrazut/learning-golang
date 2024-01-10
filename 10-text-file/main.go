package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Get Args
	filename := os.Args[1]

	if filename == "" {
		fmt.Println("Missing filename.")
	}

	// Read file
	fmt.Println("Opening file:", filename)
	buf, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	// Print
	io.Copy(os.Stdout, buf)
}
