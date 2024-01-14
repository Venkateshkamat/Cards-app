package main

//create a new type deck which is slice of strings

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string //creating something of type []string (array of strings) kind of like


func newDeck() deck {
	cards := deck{}                                                //creating object like instance of type deck
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"} //array of strings
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { // underscore for ignoring the value
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// argument function which needs args like cards(deck,size)
func deal(d deck, handSize int) (deck, deck) { //returns two values both of type deck
	return d[:handSize], d[handSize:] //return by seperating them by coma
}

// reciever fucntion that can be called cards.print
func (d deck) print() {
	fmt.Println("----------------")
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {

	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())		//creates random int64 using time now

	r := rand.New(source)			//created random seed using int64
	for i := range d {
		newPosition := r.Intn(len(d) - 1)		

		d[i], d[newPosition] = d[newPosition], d[i]		//swaps values
	}

}
