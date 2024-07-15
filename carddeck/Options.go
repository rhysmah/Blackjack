package carddeck

import "fmt"

// Options defines configuration options for a deck of cards.
//
// Fields:
//
//	sortFunc   func([]Card)           - Function to sort the deck of cards.
//	shuffle    bool                   - Flag indicating if the deck should be shuffled.
//	filterCard func(card Card) bool   - Function to filter out specific cards from the deck.
//	numJokers  int                    - Number of jokers to include in the deck.
//	numDecks   int                    - Number of decks to combine.
type Options struct {
	sortFunc   func([]Card)
	shuffle    bool
	filterCard func(card Card) bool
	numJokers  int
	numDecks   int
}

// OptionsFunc defines a function type that modifies the Options struct.
//
// Parameters:
//
//	deckOpts *Options - Pointer to an Options struct to be modified.
//
// Returns:
//
//	error - An error if the modification is invalid, otherwise nil.
type OptionsFunc func(deckOpts *Options) error

// WithSort sets the sorting function for the deck of cards.
//
// Parameters:
//
//	sortFunc func([]Card) - Function to sort the deck of cards.
//
// Returns:
//
//	OptionsFunc - A function that sets the sorting function in the Options struct.
func WithSort(sortFunc func([]Card)) OptionsFunc {
	return func(deckOpts *Options) error {
		deckOpts.sortFunc = sortFunc
		return nil
	}
}

// WithShuffle sets the shuffle flag to true.
//
// Returns:
//
//	OptionsFunc - A function that sets the shuffle flag in the Options struct.
func WithShuffle() OptionsFunc {
	return func(deckOpts *Options) error {
		deckOpts.shuffle = true
		return nil
	}
}

// WithFilteredCards sets the filtering function for the deck of cards.
//
// Parameters:
//
//	filterFunc func(card Card) bool - Function to filter out specific cards.
//
// Returns:
//
//	OptionsFunc - A function that sets the filtering function in the Options struct.
func WithFilteredCards(filterFunc func(card Card) bool) OptionsFunc {
	return func(deckOpts *Options) error {
		deckOpts.filterCard = filterFunc
		return nil
	}
}

// WithJokers sets the number of jokers to include in the deck.
//
// Parameters:
//
//	n int - Number of jokers to include.
//
// Returns:
//
//	OptionsFunc - A function that sets the number of jokers in the Options struct.
//	error       - An error if the number of jokers is negative.
func WithJokers(n int) OptionsFunc {
	return func(deckOpts *Options) error {
		if n < 0 {
			return fmt.Errorf("number of jokers cannot be negative: %d", n)
		}
		deckOpts.numJokers = n
		return nil
	}
}

// WithMultipleDecks sets the number of decks to combine.
//
// Parameters:
//
//	n int - Number of decks to combine.
//
// Returns:
//
//	OptionsFunc - A function that sets the number of decks in the Options struct.
//	error       - An error if the number of decks is negative.
func WithMultipleDecks(n int) OptionsFunc {
	return func(deckOpts *Options) error {
		if n < 0 {
			return fmt.Errorf("number of decks cannot be negative: %d", n)
		}
		deckOpts.numDecks = n
		return nil
	}
}
