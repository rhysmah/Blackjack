package hand

import (
	"blackjack/deck"
	"fmt"
)

// Represents a hand of cards in a Blackjack game
type Hand struct {
	Cards  []deck.Card
	Points int
}

// Creates a new hand with no cards and 0 points
func New() Hand {
	cards := make([]deck.Card, 0, 9)
	points := 0
	return Hand{Cards: cards, Points: points}
}

// Adds a card from a deck of cards to the hand; removes
// the card from the deck to avoid repetition
func (h *Hand) AddCard(deckOfCards *[]deck.Card) error {
	if len(*deckOfCards) == 0 {
		return fmt.Errorf("the deck is empty")
	}

	h.Cards = append(h.Cards, (*deckOfCards)[0])
	*deckOfCards = (*deckOfCards)[1:]
	return nil
}

// Updates the score of the hand
// Aces count as either 1 or 11 based on the value of hand; their
// value is determined last, after the other cards are account for
func (h *Hand) UpdateScore() {
	h.Points = 0
	acesCount := 0

	for _, card := range h.Cards {
		if card.Value == deck.Ace {
			acesCount++
		} else {
			h.Points += int(card.Value)
		}
	}

	for i := 0; i < acesCount; i++ {
		if h.Points > 10 {
			h.Points += 1
		} else {
			h.Points += 11
		}
	}
}
