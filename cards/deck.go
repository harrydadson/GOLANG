package main

import (
	"time"
	"math/rand"
	"os"
	"fmt"
	"io/ioutil"
	"strings"
)

// create a new type of 'deck', which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"} // regular list
	cardValues := []string{"Ace", "Two", "three", "Four"}

	for _, suit := range cardSuits {				// make combination of decks
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {       // added a receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // function takes 2 args(deck, handSize),return 2 values of type deck
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option1 - log the error and return a call to newDeck()
		// option2 - log the error and entirely quit the program, go with option2
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), "," )
	return deck(s) // Ace of spades, Two of Spades, Three of Spades
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)

		d[i], d[newPos] = d[newPos], d[i]
	}
}
