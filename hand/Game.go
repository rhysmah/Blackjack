package hand

import (
	"blackjack/carddeck"
	"fmt"
)

type Game struct {
	DeckOfCards []carddeck.Card
	Player      PlayerHand
	Dealer      DealerHand
}

func NewGame() (*Game, error) {
	deckOfCards, err := carddeck.New(carddeck.WithMultipleDecks(1), carddeck.WithShuffle())
	if err != nil {
		return nil, err
	}

	player, err := NewPlayer(WithStartingAmount(100.00))
	if err != nil {
		return nil, fmt.Errorf("error creating player: %w", err)
	}

	dealer, err := NewDealer(WithStartingAmount(100.00))
	if err != nil {
		return nil, fmt.Errorf("error creating dealer: %w", err)
	}

	game := &Game{
		DeckOfCards: deckOfCards,
		Player:      player,
		Dealer:      dealer,
	}

	for i := 0; i < 2; i++ {
		game.Player.AddCard(&deckOfCards)
		game.Dealer.AddCard(&deckOfCards)
	}

	return game, nil
}
