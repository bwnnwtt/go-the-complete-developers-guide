package main

import "fmt"

func main() {

	// Initialize an empty map with var
	// var colors map[string]string

	// Initialize a map with make
	// colors := make(map[string]string)

	// Initialize a map with key value pairs
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffee",
	}

	// create or update a key-value pair
	colors["white"] = "#ffffff"

	// delete a key-value pair
	delete(colors, "white")

	printMap(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println("Hex code for", k, "is", v)
	}
}
