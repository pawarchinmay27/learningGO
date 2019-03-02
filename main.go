package main

func main() {
	cards := newDeck()
	/*cards.print()
	hand,remainingDeck:=deal(cards,5)
	hand.print()
	remainingDeck.print()
	*/
	cards:=newDeck()
	fmt.Println(cards.toString())
}
func getNewCard() string {
	return "Ace Of Spades"
}

