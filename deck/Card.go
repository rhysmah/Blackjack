package deck

import (
	"fmt"
)

type Suit int  // Suit of card
type Value int // Face value of card

type Card struct {
	Suit  Suit
	Value Value
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

// Map converts card suit int to string for readability
var suits = map[Suit]string{
	0: "Spades",
	1: "Diamonds",
	2: "Clubs",
	3: "Hearts",
	4: "Joker",
}

// Map converts card value int to string for readability
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

// String() is a special method called whenever a print function is used
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
