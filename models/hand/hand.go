package hand

import (
	"pokerTools/models/card"
	"sort"
)

// Hand represents a poker hand, a set of five cards.
type Hand struct {
	Cards [5]card.Card
}

// NewHand creates a new Hand with sorted cards by value.
func NewHand(cards [5]card.Card) Hand {
	sort.Slice(cards[:], func(i, j int) bool {
		return cards[i].Value > cards[j].Value
	})
	return Hand{Cards: cards}
}
