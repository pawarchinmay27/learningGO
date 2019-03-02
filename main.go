package main
import _"fmt"
func main() {
	cards:=newDeck()
	cards.saveToFile("my_cards")
}
func getNewCard() string {
	return "Ace Of Spades"
}

