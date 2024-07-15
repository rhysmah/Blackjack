package carddeck

import (
	"fmt"
)

// The suit of card: Spades, Diamonds, Clubs, Hearts, and Jokers
type Suit int

// The value of a card: Joker (0) to King (13)
type Value int

// Card represents a playing card, comprised of a suit and a value
type Card struct {
	Suit  Suit
	Value Value
}

const (
	Spades   Suit = iota // value 0
	Diamonds             // value 1
	Clubs                // value 2
	Hearts               // value 3
	Joker                // value 4
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

// suits maps a Suit, an integer, to a string
// representation of a card's suit.
var suits = map[Suit]string{
	0: "Spades",
	1: "Diamonds",
	2: "Clubs",
	3: "Hearts",
	4: "Joker",
}

// values maps a Value, an integer, to a string
// representation of a card's value.
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

// String() is a special method called whenever a print function is used.
// It prints a simple representation of a playing card, e.g., "Ace of Hearts"
func (c Card) String() string {
	switch {
	case c.Suit == Joker:
		return fmt.Sprint(suits[c.Suit])
	case c.Value < Ace || c.Value > King:
		return fmt.Sprintf("invalid card value: %d", c.Value)
	case c.Suit < Spades || c.Suit > Hearts:
		return fmt.Sprintf("invalid card suit: %d", c.Suit)
	default:
		return fmt.Sprintf("%s of %s", values[c.Value], suits[c.Suit])
	}
}
