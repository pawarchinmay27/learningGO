package main
import _"fmt"
func main() {
	cards:=newDeck()
	cards.print()
	cards.shuffle()
	cards.print()
}
func getNewCard() string {
	return "Ace Of Spades"
}

