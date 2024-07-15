package carddeck

import (
	"log"
	"math/rand"
)

// generateDeck generates a deck of cards.
//
// This function creates one or more decks of standard playing cards based on the
// provided configuration. Each deck contains cards from Ace to King in all four suits.
// It allows for filtering out specific cards as defined by the filterCard option.
//
// Parameters:
//
//	config *Options - A pointer to an Options struct containing the configuration
//	                  for generating the deck. The configuration includes the number
//	                  of decks to generate and any card filtering criteria.
//
// Returns:
//
//	[]Card - A slice of Card objects representing the generated deck(s) of cards.
//
// Configuration Options:
//   - numDecks: The number of decks to generate. Each deck contains 52 cards.
//   - filterCard: A function that takes a Card and returns a boolean indicating
//     if the card should be excluded from the deck. If nil, no cards are filtered out.
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

// New creates a new deck of cards with specific configurations.
//
// This function generates a deck of cards based on the provided options. It allows
// customization such as the number of decks, the inclusion of jokers, shuffling,
// filtering specific cards, and sorting.
//
// Parameters:
//
//	opts ...OptionsFunc - Variadic parameter representing a list of option functions
//	                      that modify the deck configuration.
//
// Returns:
//
//	[]Card - A slice of Card objects representing the newly created deck of cards.
//	error  - An error value if any of the option functions return an error.
//
// Available Options Functions:
//   - WithJokers(n int): Adds 'n' jokers to the deck.
//   - WithShuffle(shuffle bool): Shuffles the deck if true.
//   - WithDecks(n int): Creates a deck of cards composed of 'n' decks of cards.
//   - WithSort(sortFunc func([]Card)): Sorts the deck using the provided sort function.
//   - WithFilter(filterFunc func(Card) bool): Filters out cards that do not meet the criteria.
func New(opts ...OptionsFunc) ([]Card, error) {
	config := &Options{
		sortFunc:   nil,
		shuffle:    false,
		filterCard: nil,
		numJokers:  0,
		numDecks:   1,
	}

	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
		opt(config)
	}

	deckOfCards := generateDeck(config)

	for i := 0; i < config.numJokers; i++ {
		deckOfCards = append(deckOfCards, Card{Suit: Joker, Value: JokerValue})
	}
	if config.shuffle {
		shuffle(deckOfCards)
	}
	if config.sortFunc != nil {
		config.sortFunc(deckOfCards)
	}
	log.Println("Successfully created deck of cards")
	return deckOfCards, nil
}

// shuffle shuffles a deck of cards.
//
// This function takes a slice of Card objects and randomly shuffles their order
// uring the rand.Shuffle function.
//
// Parameters:
//
//	deckOfCards []Card - A slice representing the deck of cards to be shuffled.
func shuffle(deckOfCards []Card) {
	rand.Shuffle(len(deckOfCards), func(i, j int) {
		deckOfCards[i], deckOfCards[j] = deckOfCards[j], deckOfCards[i]
	})
}
