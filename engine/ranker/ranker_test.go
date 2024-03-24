package ranker

import (
	"pokerTools/models/card"
	"pokerTools/models/hand"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	straightFlush = hand.NewHand(
		[5]card.Card{
			{
				Value: card.Three,
				Color: card.Spade,
			},
			{
				Value: card.Seven,
				Color: card.Spade,
			},
			{
				Value: card.Four,
				Color: card.Spade,
			},
			{
				Value: card.Five,
				Color: card.Spade,
			},
			{
				Value: card.Six,
				Color: card.Spade,
			},
		})

	fourOfAKind = hand.NewHand(
		[5]card.Card{
			{
				Value: card.Three,
				Color: card.Spade,
			},
			{
				Value: card.Three,
				Color: card.Heart,
			},
			{
				Value: card.Three,
				Color: card.Diamond,
			},
			{
				Value: card.Six,
				Color: card.Spade,
			},
			{
				Value: card.Three,
				Color: card.Club,
			},
		})

	full = hand.NewHand(
		[5]card.Card{
			{
				Value: card.Three,
				Color: card.Spade,
			},
			{
				Value: card.Three,
				Color: card.Heart,
			},
			{
				Value: card.Five,
				Color: card.Heart,
			},
			{
				Value: card.Five,
				Color: card.Spade,
			},
			{
				Value: card.Three,
				Color: card.Club,
			},
		})

	flush = hand.NewHand(
		[5]card.Card{
			{
				Value: card.Two,
				Color: card.Spade,
			},
			{
				Value: card.Three,
				Color: card.Spade,
			},
			{
				Value: card.Jack,
				Color: card.Spade,
			},
			{
				Value: card.Five,
				Color: card.Spade,
			},
			{
				Value: card.Ten,
				Color: card.Spade,
			},
		})

	straight = hand.NewHand(
		[5]card.Card{
			{
				Value: card.Two,
				Color: card.Spade,
			},
			{
				Value: card.Three,
				Color: card.Spade,
			},
			{
				Value: card.Four,
				Color: card.Spade,
			},
			{
				Value: card.Five,
				Color: card.Spade,
			},
			{
				Value: card.Six,
				Color: card.Heart,
			},
		})

	threeOfAKind = hand.NewHand(
		[5]card.Card{
			{
				Value: card.Two,
				Color: card.Spade,
			},
			{
				Value: card.Ace,
				Color: card.Club,
			},
			{
				Value: card.Two,
				Color: card.Club,
			},
			{
				Value: card.Jack,
				Color: card.Spade,
			},
			{
				Value: card.Two,
				Color: card.Heart,
			},
		})

	doublePair = hand.NewHand(
		[5]card.Card{
			{
				Value: card.Two,
				Color: card.Spade,
			},
			{
				Value: card.Ace,
				Color: card.Club,
			},
			{
				Value: card.Two,
				Color: card.Club,
			},
			{
				Value: card.Ace,
				Color: card.Spade,
			},
			{
				Value: card.Queen,
				Color: card.Heart,
			},
		})

	pair = hand.NewHand(
		[5]card.Card{
			{
				Value: card.Two,
				Color: card.Spade,
			},
			{
				Value: card.Ace,
				Color: card.Club,
			},
			{
				Value: card.Six,
				Color: card.Club,
			},
			{
				Value: card.Ace,
				Color: card.Spade,
			},
			{
				Value: card.King,
				Color: card.Heart,
			},
		})

	highCard = hand.NewHand(
		[5]card.Card{
			{
				Value: card.Two,
				Color: card.Spade,
			},
			{
				Value: card.Five,
				Color: card.Club,
			},
			{
				Value: card.Six,
				Color: card.Club,
			},
			{
				Value: card.Ace,
				Color: card.Spade,
			},
			{
				Value: card.King,
				Color: card.Heart,
			},
		})
)

func TestComputeStrength(t *testing.T) {
	t.Run("high card", func(t *testing.T) {
		assert.Equal(t, Strength("0ed652"), ComputeStrength(highCard))
	})

	t.Run("pair", func(t *testing.T) {
		assert.Equal(t, Strength("1ed62"), ComputeStrength(pair))
	})

	t.Run("double pair", func(t *testing.T) {
		assert.Equal(t, Strength("2e2c"), ComputeStrength(doublePair))
	})

	t.Run("three of a kind", func(t *testing.T) {
		assert.Equal(t, Strength("32eb"), ComputeStrength(threeOfAKind))
	})

	t.Run("straight", func(t *testing.T) {
		assert.Equal(t, Strength("46"), ComputeStrength(straight))
	})

	t.Run("flush", func(t *testing.T) {
		assert.Equal(t, Strength("5ba532"), ComputeStrength(flush))
	})

	t.Run("full", func(t *testing.T) {
		assert.Equal(t, Strength("635"), ComputeStrength(full))
	})

	t.Run("four of a kind", func(t *testing.T) {
		assert.Equal(t, Strength("736"), ComputeStrength(fourOfAKind))
	})

	t.Run("straight flush", func(t *testing.T) {
		assert.Equal(t, Strength("87"), ComputeStrength(straightFlush))
	})

	t.Run("order", func(t *testing.T) {
		expectedOrder := []Strength{
			ComputeStrength(highCard),
			ComputeStrength(pair),
			ComputeStrength(doublePair),
			ComputeStrength(threeOfAKind),
			ComputeStrength(straight),
			ComputeStrength(flush),
			ComputeStrength(full),
			ComputeStrength(fourOfAKind),
			ComputeStrength(straightFlush),
		}
		assert.True(t, slices.IsSorted(expectedOrder))
	})

}
