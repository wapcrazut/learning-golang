package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":  "#ff0000",
		"blue": "#0000ff",
	}

	delete(colors, "blue")

	colors["white"] = "#ffffff"

	printColorMap(colors)
}

func printColorMap(m map[string]string) {
	for c, h := range m {
		fmt.Println(c, "hex code is:", h)
	}
}
