package deck

import (
	"pokerTools/models/card"
)

type Deck struct {
	Cards [7]card.Card
}

func (d Deck) makePossibleHands() (hands [21]hand) {
	for handIndex, indexes := range fiveOutOfSevenCombinationsIndexes {
		hands[handIndex] = d.makeSingleHandFromIndexes(indexes)
	}

	return hands
}

func (d Deck) makeSingleHandFromIndexes(cardIndexes [5]int) hand {
	cards := [5]card.Card{}
	for k, idx := range cardIndexes {
		cards[k] = d.Cards[idx]
	}

	return hand{cards: cards}
}
