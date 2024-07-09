package hand

import (
	"blackjack/deck"
	"fmt"
)

// Add configOpts via functional options for clearer configurations

type BasicHand struct {
	Cards  []deck.Card
	Points int
}

type Player struct {
	BasicHand
}

type Dealer struct {
	BasicHand
}

func NewPlayer() Player {
	return Player{
		BasicHand: BasicHand{
			Cards:  make([]deck.Card, 0, 9),
			Points: 0,
		},
	}
}

func NewDealer() Dealer {
	return Dealer{
		BasicHand: BasicHand{
			Cards:  make([]deck.Card, 0, 9),
			Points: 0,
		},
	}
}

// Adds a card from a deck of cards to the hand; removes
// the card from the deck to avoid repetition
func (h *BasicHand) AddCard(deckOfCards *[]deck.Card) error {
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
func (h *BasicHand) UpdateScore() {
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

// Displays the player hand. All cards are always revealed.
func (p *Player) DisplayHand() {
	fmt.Println("Player's Cards:")
	for _, card := range p.Cards {
		fmt.Printf("%s ", card)
	}
	fmt.Println()
}

// Displays the dealer's hand. The second card remains hidden
// until the final hand, where all the dealer's cards are revealed
func (d *Dealer) DisplayHand(isFinalHand bool) {
	fmt.Println("Dealer's Cards:")
	for i := 0; i < len(d.Cards); i++ {
		if i == 1 && !isFinalHand {
			fmt.Printf("Hidden Card\n")
		} else {
			fmt.Printf("%s ", d.Cards[i])
		}
	}
	fmt.Println()
}
