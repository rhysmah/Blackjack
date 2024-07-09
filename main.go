package main

import (
	"blackjack/deck"
	"blackjack/hand"
	"fmt"
)

func main() {

	deck := deck.New(deck.WithSort(deck.SortByValue))
	hand := hand.New()

	fmt.Println(deck)
	fmt.Println(hand.Points)

	hand.AddCard(&deck)
	hand.AddCard(&deck)
	hand.UpdateScore()

	fmt.Println(hand.Cards)
	fmt.Println(hand.Points)
}
