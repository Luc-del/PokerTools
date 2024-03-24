package hand

import (
	"pokerTools/models/card"
	"sort"
)

type Hand struct {
	cards [5]card.Card
}

func NewHand(cards [5]card.Card) Hand {
	sort.Slice(cards[:], func(i, j int) bool {
		return cards[i].Value > cards[j].Value
	})
	return Hand{cards: cards}
}
