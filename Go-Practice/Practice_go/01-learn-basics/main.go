package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// var name type = expression
	var integer int = 10
	integer1, string := 1, "string"

	fmt.Println(integer)
	fmt.Println(integer1, string)

	// name := expression
	boolean := true
	fmt.Println(boolean)

	// Pointers
	x := 1
	p := &x

	fmt.Println(*p)

	// Declaration types
	// fahrenheit & Celsius
	type fahrenheit int
	type celsius int

	var f fahrenheit = 32
	var c celsius = 0

	c = celsius((f - 32) * 5 / 9)

	fmt.Println(f, c)


}