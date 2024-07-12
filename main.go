package main

import (
	"blackjack/deck"
	"blackjack/hand"
	"fmt"
	"strings"
)

type CommandFunc func(deck *[]deck.Card, player *hand.Player)

// Player adds a card to their hand
func hit(deck *[]deck.Card, player *hand.Player) {
	player.AddCard(deck)
}

// Play stays; nothing happens
func stay(deck *[]deck.Card, player *hand.Player) {
}

// Gets the user's input; returns a string trimmed of whitespace
func getUserOption() string {
	var userChoice string
	fmt.Println("Do you hit (1) or stay (2)?")
	fmt.Scanln(&userChoice)
	return strings.TrimSpace(userChoice)
}

// Gameplay loop that asks the player to hit or stay.
// If the player busts (> 21) or gets blackjack (== 21), the loop ends.
// If the player stays, the loop ends
func play(deck *[]deck.Card, dealer *hand.Dealer, player *hand.Player) (bool, bool) {
	playerOptions := map[string]CommandFunc{
		"1": hit,
		"2": stay,
	}

	for {
		player.UpdateScore()

		if player.Points > 21 {
			return true, false
		}
		if player.Points == 21 {
			return false, true
		}

		player.DisplayHand()
		dealer.DisplayHand()

		userChoice := getUserOption()
		if cmd, ok := playerOptions[userChoice]; ok {
			cmd(deck, player)
			if userChoice == "2" {
				break
			}
		}
		fmt.Println("Invalid selection. Select 1 (hit) or 2 (stay)")
	}
	return false, false
}

// Deals the player and dealer two cards each
func dealInitialHands(deck *[]deck.Card, playerHand, dealerHand *hand.BasicHand) {
	for i := 0; i < 2; i++ {
		playerHand.AddCard(deck)
		dealerHand.AddCard(deck)
	}
}

// Displays the dealer's and player's hands and checks for a winner
func endGame(bust, blackjack bool, player hand.Player, dealer hand.Dealer) {
	player.DisplayHand()
	dealer.DisplayHand(hand.IsFinalHand())

	if bust {
		fmt.Println("You bust! Dealer wins.")
	} else if blackjack {
		fmt.Println("You got blackjack! You win.")
	} else if player.Points > dealer.Points {
		fmt.Println("You got the higher score. You win.")
	} else if player.Points < dealer.Points {
		fmt.Println("The dealer got the higher score. You lose.")
	} else {
		fmt.Println("Tie! Nobody wins.")
	}
}

func main() {
	deck, err := deck.New(deck.WithShuffle())
	if err != nil {
		panic(err)
	}
	player := hand.NewPlayer()
	dealer := hand.NewDealer()

	dealInitialHands(&deck, &player.BasicHand, &dealer.BasicHand)

	// Gameplay loop
	bust, blackjack := play(&deck, &dealer, &player)

	endGame(bust, blackjack, player, dealer)
}
