package ranker_test

import (
	"pokerTools/engine/ranker"
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
		assert.Equal(t, ranker.Strength("0ed652"), ranker.ComputeStrength(highCard))
	})

	t.Run("pair", func(t *testing.T) {
		assert.Equal(t, ranker.Strength("1ed62"), ranker.ComputeStrength(pair))
	})

	t.Run("double pair", func(t *testing.T) {
		assert.Equal(t, ranker.Strength("2e2c"), ranker.ComputeStrength(doublePair))
	})

	t.Run("three of a kind", func(t *testing.T) {
		assert.Equal(t, ranker.Strength("32eb"), ranker.ComputeStrength(threeOfAKind))
	})

	t.Run("straight", func(t *testing.T) {
		assert.Equal(t, ranker.Strength("46"), ranker.ComputeStrength(straight))
	})

	t.Run("flush", func(t *testing.T) {
		assert.Equal(t, ranker.Strength("5ba532"), ranker.ComputeStrength(flush))
	})

	t.Run("full", func(t *testing.T) {
		assert.Equal(t, ranker.Strength("635"), ranker.ComputeStrength(full))
	})

	t.Run("four of a kind", func(t *testing.T) {
		assert.Equal(t, ranker.Strength("736"), ranker.ComputeStrength(fourOfAKind))
	})

	t.Run("straight flush", func(t *testing.T) {
		assert.Equal(t, ranker.Strength("87"), ranker.ComputeStrength(straightFlush))
	})

	t.Run("order", func(t *testing.T) {
		expectedOrder := []ranker.Strength{
			ranker.ComputeStrength(highCard),
			ranker.ComputeStrength(pair),
			ranker.ComputeStrength(doublePair),
			ranker.ComputeStrength(threeOfAKind),
			ranker.ComputeStrength(straight),
			ranker.ComputeStrength(flush),
			ranker.ComputeStrength(full),
			ranker.ComputeStrength(fourOfAKind),
			ranker.ComputeStrength(straightFlush),
		}
		assert.True(t, slices.IsSorted(expectedOrder))
	})
}
