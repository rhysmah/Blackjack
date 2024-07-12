package deck

import (
	"log"
	"math/rand"
)

// Generates a deck of cards
func generateDeck(config *Options) []Card {
	var deck []Card
	for i := 0; i < config.numDecks; i++ {
		for suit := Spades; suit <= Hearts; suit++ {
			for value := Ace; value <= King; value++ {
				card := Card{Suit: suit, Value: value}

				if config.filterCard == nil || !config.filterCard(card) {
					deck = append(deck, card)
				}
			}
		}
	}
	return deck
}

// Creates a new deck of cards with specific configurations
func New(opts ...OptionsFunc) ([]Card, error) {
	defaultConfig := &Options{
		sortFunc:   nil,
		shuffle:    false,
		filterCard: nil,
		numJokers:  0,
		numDecks:   1,
	}

	for _, opt := range opts {
		if err := opt(defaultConfig); err != nil {
			return nil, err
		}
		opt(defaultConfig)
	}

	deckOfCards := generateDeck(defaultConfig)

	for i := 0; i < defaultConfig.numJokers; i++ {
		deckOfCards = append(deckOfCards, Card{Suit: JokerSuit, Value: JokerValue})
	}
	if defaultConfig.shuffle {
		shuffle(deckOfCards)
	}
	if defaultConfig.sortFunc != nil {
		defaultConfig.sortFunc(deckOfCards)
	}
	log.Println("Successfully created deck of cards")
	return deckOfCards, nil
}

// Shuffles a deck of cards
func shuffle(deckOfCards []Card) {
	rand.Shuffle(len(deckOfCards), func(i, j int) {
		deckOfCards[i], deckOfCards[j] = deckOfCards[j], deckOfCards[i]
	})
}
