package main

import ("fmt"
		"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"spades", "hearts", "diamonds", "clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+"of"+cardSuit)
		}
	}
	return cards
}


//this is a reciever function the 
//in (d deck) the d here holds the reference of the calling object
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck,handSize int) (deck,deck) {
	return d[:handSize],d[handSize:]
}

func (d deck) toString()(string){
	//type cast deck to slice of string []string(d)
	//joins  every string in slice with the give sep and returns a single string
	return strings.Join([]string(d),",")

}