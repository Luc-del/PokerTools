package deck

import (
	"github.com/stretchr/testify/assert"
	"pokerTools/models/card"
	"pokerTools/models/combination"
	"testing"
)

func Test_hand_computeHandValue(t *testing.T) {
	tests := []struct {
		name     string
		hand     hand
		expected combination.HandValue
	}{
		{
			name: "royal straight flush",
			hand: newHandDSL().withColors(card.Spade).
				withValues(card.Ace, card.Jack, card.Queen, card.King, card.Ten).build(),
			expected: combination.HandValue{
				Combination: combination.StraightFlush,
				Comparator:  combination.StraightFlushComparator{HighestValue: card.Ace},
			},
		},
		{
			name: "straight flush",
			hand: newHandDSL().withColors(card.Spade).
				withValues(card.Eight, card.Jack, card.Seven, card.Ten, card.Nine).build(),
			expected: combination.HandValue{
				Combination: combination.StraightFlush,
				Comparator:  combination.StraightFlushComparator{HighestValue: card.Jack},
			},
		},
		{
			name: "four of a kind",
			hand: newHandDSL().withValues(card.Eight, card.Eight, card.Eight, card.Ten, card.Eight).build(),
			expected: combination.HandValue{
				Combination: combination.FourOfAKind,
				Comparator:  combination.FourOfAKindComparator{Value: card.Eight},
			},
		},
		{
			name: "full",
			hand: newHandDSL().withValues(card.Two, card.Nine, card.Two, card.Two, card.Nine).build(),
			expected: combination.HandValue{
				Combination: combination.Full,
				Comparator: combination.FullComparator{
					ThreeOfAKindValue: card.Two,
					PairValue:         card.Nine,
				},
			},
		},
		{
			name: "flush",
			hand: newHandDSL().withColors(card.Spade).
				withValues(card.Eight, card.Ace, card.Four, card.Six, card.Nine).build(),
			expected: combination.HandValue{
				Combination: combination.Flush,
				Comparator:  combination.FlushComparator{HighestValue: card.Ace},
			},
		},
		{
			name: "straight",
			hand: newHandDSL().withColors(card.Club, card.Spade, card.Diamond, card.Club, card.Heart).
				withValues(card.Eight, card.Six, card.Seven, card.Five, card.Nine).build(),
			expected: combination.HandValue{
				Combination: combination.Straight,
				Comparator:  combination.StraightComparator{HighestValue: card.Nine},
			},
		},
		{
			name: "three of a kind",
			hand: newHandDSL().withValues(card.Queen, card.Eight, card.Eight, card.Ten, card.Eight).build(),
			expected: combination.HandValue{
				Combination: combination.ThreeOfAKind,
				Comparator:  combination.ThreeOfAKindComparator{Value: card.Eight},
			},
		},
		{
			name: "double pair",
			hand: newHandDSL().withValues(card.Queen, card.Five, card.Queen, card.Six, card.Six).build(),
			expected: combination.HandValue{
				Combination: combination.DoublePair,
				Comparator: combination.DoublePairComparator{
					PairValues:    [2]card.Value{card.Queen, card.Six},
					HighCardValue: card.Five,
				},
			},
		},
		{
			name: "pair",
			hand: newHandDSL().withValues(card.King, card.Five, card.Nine, card.Six, card.Nine).build(),
			expected: combination.HandValue{
				Combination: combination.Pair,
				Comparator: combination.PairComparator{
					PairValue:  card.Nine,
					CardValues: [3]card.Value{card.King, card.Six, card.Five},
				},
			},
		},
		{
			name: "high card",
			hand: newHandDSL().withValues(card.Six, card.Five, card.Ten, card.Three, card.Nine).build(),
			expected: combination.HandValue{
				Combination: combination.HighCard,
				Comparator: combination.HighCardComparator{
					CardValues: [5]card.Value{card.Ten, card.Nine, card.Six, card.Five, card.Three},
				},
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.hand.computeHandValue())
		})
	}
}
