package main

import "fmt"

func main() {

	colors := map[string]string{ // says all keys are string, and values are strings as well
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#fffff",
	}
/*
	colors := make(map[int]string)

	colors[10] = "ffffff"

	delete(colors, 10) // to delete
*/
	printMap(colors)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println("Hex code for", k, "is", v )
	}
}


/*

Maps												Structs
All keys must be same type							Values can be different types
All values must be the same type					Keys dont support indexing
keys are indexed, can iterate over					to represent a thing with lots of diffent properties
To represent collection of related properties		Value Type
Reference Type - map[string]string					[]string{}

*/
