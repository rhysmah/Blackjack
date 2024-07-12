package main

import (
	"blackjack/carddeck"
	"blackjack/hand"
	"fmt"
	"strings"
)

const (
	HitOption  = "1"
	StayOption = "2"
)

type CommandFunc func(deck *[]carddeck.Card, player *hand.PlayerHand)

// Player adds a card to their hand
func hit(deck *[]carddeck.Card, player *hand.PlayerHand) {
	player.AddCard(deck)
}

// Play stays; nothing happens
func stay(deck *[]carddeck.Card, player *hand.PlayerHand) {}

// Gets the user's input; returns a string trimmed of whitespace
func getUserOption() string {
	var userChoice string
	fmt.Println("Do you hit (1) or stay (2)?")
	fmt.Scanln(&userChoice)
	return strings.TrimSpace(userChoice)
}

func checkWinConditions(player *hand.PlayerHand, dealer *hand.DealerHand) (bool, bool, bool) {
	switch {
	case player.Points > 21:
		return false, true, false // player busts, dealer wins
	case dealer.Points > 21:
		return true, false, false // player wins, dealer busts
	case player.Points == 21 && len(player.Cards) == 2 && dealer.Points == 21 && len(dealer.Cards) == 2:
		return false, false, true // tie: player and dealer both get blackjack
	case player.Points == 21 && len(player.Cards) == 2:
		return true, false, false // player wins with blackjack
	case dealer.Points == 21 && len(dealer.Cards) == 2:
		return false, true, false // dealer wins with blackjack
	default:
		return false, false, false // Nobody wins yet
	}
}

func checkFinalWinConditions(player *hand.PlayerHand, dealer *hand.DealerHand) (bool, bool, bool) {
	switch {
	case player.Points > dealer.Points:
		return true, false, false // player wins
	case player.Points < dealer.Points:
		return false, true, false // dealer wins
	default:
		return false, false, true // tie: dealer and player both have non-blackjack 21
	}
}

// Plays the game
func play(deck *[]carddeck.Card, dealer *hand.DealerHand, player *hand.PlayerHand) (bool, bool, bool) {
	playerOptions := map[string]CommandFunc{
		HitOption:  hit,
		StayOption: stay,
	}

	// Gameplay loop
	for {
		player.CalculateScore()
		dealer.CalculateScore()

		playerWin, dealerWin, tie := checkWinConditions(player, dealer)
		if playerWin || dealerWin || tie {
			return playerWin, dealerWin, tie
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

	player.CalculateScore()
	dealer.CalculateScore()

	return checkFinalWinConditions(player, dealer)
}

// Displays the dealer's and player's hands and checks for a winner
func endGame(playerWins, dealerWins, tie bool, player hand.PlayerHand, dealer hand.DealerHand) {

	player.DisplayHand()
	dealer.DisplayHand(hand.IsFinalHand())

	switch {
	case playerWins:
		fmt.Println("You win!")
	case dealerWins:
		fmt.Println("Dealer wins!")
	case tie:
		fmt.Println("It's a tie!")
	}
}

func main() {
	game, err := hand.NewGame()
	if err != nil {
		panic(err)
	}
	playerWins, dealerWins, tie := play(&game.DeckOfCards, &game.Dealer, &game.Player)
	endGame(playerWins, dealerWins, tie, game.Player, game.Dealer)
}
