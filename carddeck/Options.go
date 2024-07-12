package carddeck

import "fmt"

// Configuration options for deck of cards
type Options struct {
	sortFunc   func([]Card)
	shuffle    bool
	filterCard func(card Card) bool
	numJokers  int
	numDecks   int
}

// Function that takes in a pointer to an Options struct and modifies it
// Allows for creation of a deck of cards with different configurations
type OptionsFunc func(deckOpts *Options) error

// Determines how to sort a deck of cards: by value or by suit
func WithSort(sortFunc func([]Card)) OptionsFunc {
	return func(DeckOpts *Options) error {
		DeckOpts.sortFunc = sortFunc
		return nil
	}
}

// Determines if a deck of cards should be shuffled
func WithShuffle() OptionsFunc {
	return func(deckOpts *Options) error {
		deckOpts.shuffle = true
		return nil
	}
}

// Determines if a deck of cards should be filtered: by suit, value or both
func WithFilteredCards(filterFunc func(card Card) bool) OptionsFunc {
	return func(deckOpts *Options) error {
		deckOpts.filterCard = filterFunc
		return nil
	}
}

// Determines how many jokers should be added to card deck
func WithJokers(n int) OptionsFunc {
	return func(deckOpts *Options) error {
		if n < 0 {
			return fmt.Errorf("number of jokers cannot be negative: %d", n)
		}
		deckOpts.numJokers = n
		return nil
	}
}

// Determines how many decks should be combined to create a deck of cards
func WithMultipleDecks(n int) OptionsFunc {
	return func(deckOpts *Options) error {
		if n < 0 {
			return fmt.Errorf("number of decks cannot be negative: %d", n)
		}
		deckOpts.numDecks = n
		return nil
	}
}
