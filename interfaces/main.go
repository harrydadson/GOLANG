package main

import "fmt"

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

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string { // or (englishBot)
	// very custom logic for generating english logic
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string { // or (spanishBot)
	return "Hola!"
}

/*

Interfaces are not generic types
Interfaces are implicit
Interfaces are a contract to help us manage types
Interfaces are tough (Step #1 is understand how to read them)

*/