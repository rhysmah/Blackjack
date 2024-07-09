package main

import (
	"blackjack/deck"
	"blackjack/hand"
)

// TODO: Create gameplay loop

func main() {

	deck := deck.New(deck.WithShuffle())

	playerHand := hand.New(false)
	dealerHand := hand.New(true)

	// Round 1
	playerHand.AddCard(&deck)
	dealerHand.AddCard(&deck)

	// Round 2
	playerHand.AddCard(&deck)
	dealerHand.AddCard(&deck)

	playerHand.DisplayPlayerHand()
	dealerHand.DisplayDealerHand(false)

	playerHand.DisplayPlayerHand()
	dealerHand.DisplayDealerHand(true)
}
