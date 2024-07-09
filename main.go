package main

import (
	"blackjack/deck"
	"blackjack/hand"
)

// TODO: Create gameplay loop

func main() {

	deck := deck.New(deck.WithShuffle())

	playerHand := hand.NewPlayer()
	dealerHand := hand.NewDealer()

	// Round 1
	playerHand.AddCard(&deck)
	dealerHand.AddCard(&deck)

	// Round 2
	playerHand.AddCard(&deck)
	dealerHand.AddCard(&deck)

	playerHand.DisplayHand()
	dealerHand.DisplayHand(true)
}
