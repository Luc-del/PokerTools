package hand_test

import (
	"pokerTools/models/card"
	"pokerTools/models/hand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHand(t *testing.T) {
	t.Run("should sort the hand by value", func(t *testing.T) {
		cards := [5]card.Card{
			{
				Value: card.Two,
				Color: card.Spade,
			},
			{
				Value: card.Five,
				Color: card.Club,
			},
			{
				Value: card.King,
				Color: card.Heart,
			},
			{
				Value: card.Jack,
				Color: card.Spade,
			},
			{
				Value: card.Seven,
				Color: card.Diamond,
			},
		}

		h := hand.NewHand(cards)
		for i := 0; i < 4; i++ {
			assert.True(t, h.Cards[i+1].Value <= h.Cards[i].Value)
		}
	})
}
