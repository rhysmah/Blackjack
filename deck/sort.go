package deck

import "sort"

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
