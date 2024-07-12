package hand

import (
	"blackjack/carddeck"
	"fmt"
)

type BasicHand struct {
	Cards  []carddeck.Card
	Points int
}

type PlayerHand struct {
	BasicHand
}

type DealerHand struct {
	BasicHand
}

func NewPlayer() PlayerHand {
	return PlayerHand{
		BasicHand: BasicHand{
			Cards:  make([]carddeck.Card, 0, 9),
			Points: 0,
		},
	}
}

func NewDealer() DealerHand {
	return DealerHand{
		BasicHand: BasicHand{
			Cards:  make([]carddeck.Card, 0, 9),
			Points: 0,
		},
	}
}

type ConfigDealerOpts struct {
	isFinalHand bool
}

type ConfigDealerOptsFunc func(dealerOpts *ConfigDealerOpts)

func IsFinalHand() ConfigDealerOptsFunc {
	return func(dealerOpts *ConfigDealerOpts) {
		dealerOpts.isFinalHand = true
	}
}

// Add card from deck to hand
func (h *BasicHand) AddCard(deckOfCards *[]carddeck.Card) error {
	if len(*deckOfCards) == 0 {
		return fmt.Errorf("the deck is empty")
	}
	h.Cards = append(h.Cards, (*deckOfCards)[0])
	*deckOfCards = (*deckOfCards)[1:]
	return nil
}

// Update hand score; aces count as either 1 or 11 based on hand value
func (h *BasicHand) CalculateScore() {
	h.Points = 0
	acesCount := 0

	for _, card := range h.Cards {
		switch card.Value {
		case carddeck.Ace:
			acesCount++
		case carddeck.Jack, carddeck.Queen, carddeck.King:
			h.Points += 10
		default:
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

// Displays player hand with scores
func (p *PlayerHand) DisplayHand() {
	fmt.Println("##############")
	fmt.Println("Player's Cards")
	fmt.Println("##############")
	for _, card := range p.Cards {
		fmt.Printf("%s\n", card)
	}
	p.CalculateScore()
	fmt.Printf("Score: %d\n", p.Points)
	fmt.Println()
}

// Displays dealer's hand. Second card remains hidden
// until the final hand, where all cards are revealed
func (d *DealerHand) DisplayHand(opts ...ConfigDealerOptsFunc) {
	defConfig := &ConfigDealerOpts{
		isFinalHand: false,
	}
	for _, opt := range opts {
		opt(defConfig)
	}

	fmt.Println("##############")
	fmt.Println("Dealer's Cards")
	fmt.Println("##############")
	for i := 0; i < len(d.Cards); i++ {
		if i == 1 && !defConfig.isFinalHand {
			fmt.Printf("Hidden Card\n")
		} else {
			fmt.Printf("%s\n", d.Cards[i])
		}
	}
	d.CalculateScore()
	if defConfig.isFinalHand {
		fmt.Printf("Score: %d\n", d.Points)
	}
	fmt.Println()
}
