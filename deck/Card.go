package deck

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
)

type Suit int
type Value int

type Card struct {
	Suit  Suit
	Value Value
}

type BySuit []Card

func (a BySuit) Len() int           { return len(a) }
func (a BySuit) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySuit) Less(i, j int) bool { return a[i].Suit < a[j].Suit }

func SortBySuit(deckOfCards []Card) {
	sort.Sort(BySuit(deckOfCards))
}

type ByValue []Card

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].Value < a[j].Value }

func SortByValue(deckOfCards []Card) {
	sort.Sort(ByValue(deckOfCards))
}

const (
	Spades    Suit = iota // value 0
	Diamonds              // value 1
	Clubs                 // value 2
	Hearts                // value 3
	JokerSuit             // value 4
)

const (
	JokerValue Value = iota // value 0
	Ace                     // value 1
	Two                     // value 2
	Three                   // value 3
	Four                    // value 4
	Five                    // value 5
	Six                     // value 6
	Seven                   // value 7
	Eight                   // value 8
	Nine                    // value 9
	Ten                     // value 10
	Jack                    // value 11
	Queen                   // value 12
	King                    // value 13
)

// Maps for the string representation of the suit and value
// Allows for readable output when printing a card, e.g. "Ace of Spades" instead of {0, 0}
var suits = map[Suit]string{
	0: "Spades",
	1: "Diamonds",
	2: "Clubs",
	3: "Hearts",
	4: "Joker",
}

var values = map[Value]string{
	1:  "Ace",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Jack",
	12: "Queen",
	13: "King",
}

// String() is a special method that's called whenever a print function is used
func (c Card) String() string {
	switch {
	case c.Suit == JokerSuit:
		return fmt.Sprint(suits[c.Suit])
	case c.Value < Ace || c.Value > King:
		return fmt.Sprintf("invalid card value: %d", c.Value)
	case c.Suit < Spades || c.Suit > Hearts:
		return fmt.Sprintf("invalid card suit: %d", c.Suit)
	default:
		return fmt.Sprintf("%s of %s", values[c.Value], suits[c.Suit])
	}
}

// Configuration options for a deck of cards
type Options struct {
	sortFunc   func([]Card)
	shuffle    bool
	filterCard func(card Card) bool
	numJokers  int
	numDecks   int
}

// Function that takes in a pointer to a DeckOptions struct and modifies it
// This allows for the creation of a deck of cards with different configurations
type OptionsFunc func(deckOpts *Options) error

// Determines how a deck of cards should be sorted
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

// Determines if a deck of cards should be filtered
// Decks can be filtered by suit, value, or both
func WithFilteredCards(filterFunc func(card Card) bool) OptionsFunc {
	return func(deckOpts *Options) error {
		deckOpts.filterCard = filterFunc
		return nil
	}
}

// Determines how many jokers should be added to a deck of cards
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

// Creates a complete deck of cards
func New(opts ...OptionsFunc) []Card {
	defaultConfig := &Options{
		sortFunc:   nil,
		shuffle:    false,
		filterCard: nil,
		numJokers:  0,
		numDecks:   1,
	}

	for _, opt := range opts {
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
	return deckOfCards
}

func shuffle(deckOfCards []Card) {
	rand.Shuffle(len(deckOfCards), func(i, j int) {
		deckOfCards[i], deckOfCards[j] = deckOfCards[j], deckOfCards[i]
	})
}
