package carddeck

import "sort"

// BySuit is a type that implements the sort.Interface for a slice of Card objects, sorting by suit.
type BySuit []Card

func (a BySuit) Len() int           { return len(a) }
func (a BySuit) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySuit) Less(i, j int) bool { return a[i].Suit < a[j].Suit }

// SortBySuit sorts a slice of Card objects by suit.
//
// Parameters:
//
//	deckOfCards []Card - A slice of Card objects to be sorted.
func SortBySuit(deckOfCards []Card) {
	sort.Sort(BySuit(deckOfCards))
}

// ByValue is a type that implements the sort.Interface for a slice of Card objects, sorting by value.
type ByValue []Card

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].Value < a[j].Value }

// SortByValue sorts a slice of Card objects by value.
//
// Parameters:
//
//	deckOfCards []Card - A slice of Card objects to be sorted.
func SortByValue(deckOfCards []Card) {
	sort.Sort(ByValue(deckOfCards))
}
