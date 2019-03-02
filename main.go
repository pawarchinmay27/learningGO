package main
import _"fmt"
func main() {
	cards:=newDeckFromFile("my_cards")
	cards.print()
}
func getNewCard() string {
	return "Ace Of Spades"
}

