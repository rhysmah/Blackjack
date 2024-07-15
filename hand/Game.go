package hand

import "blackjack/carddeck"

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

	game := &Game{
		DeckOfCards: deckOfCards,
		Player:      NewPlayer(),
		Dealer:      NewDealer(),
	}

	for i := 0; i < 2; i++ {
		game.Player.AddCard(&deckOfCards)
		game.Dealer.AddCard(&deckOfCards)
	}

	return game, nil
}
