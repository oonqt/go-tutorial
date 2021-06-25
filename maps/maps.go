package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// var colors map[string]string

	// colors := make(map[int]string)

	// colors[10] = "#ffffff"
	// colors[9] = "#4bf745"

	// delete(colors, 9)

	// fmt.Print(colors);

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for", color, "is", hex)
	}
}