package deck

import (
	"pokerTools/models/card"
	"pokerTools/models/combination"
	"sort"
)

type hand struct {
	cards [5]card.Card
}

func (h hand) Values() (values [5]card.Value) {
	for idx, c := range h.cards {
		values[idx] = c.Value
	}

	return values
}

func (h hand) computeHandValue() (res combination.HandValue) {
	sort.Slice(h.cards[:], func(i, j int) bool {
		return h.cards[i].Value < h.cards[j].Value
	})

	switch {
	case h.isStraightFlush(&res):
	case h.isFourOfAKind(&res):
	case h.isFull(&res):
	case h.isFlush(&res):
	case h.isStraight(&res):
	case h.isThreeOfAKind(&res):
	case h.isDoublePair(&res):
	case h.isPair(&res):
	default:
		h.isHighCard(&res)
	}

	return res
}

func (h hand) isStraightFlush(hv *combination.HandValue) bool {
	compareCard := h.cards[0]
	for idx := 1; idx < 5; idx++ {
		tmpCard := h.cards[idx]
		if tmpCard.Color != compareCard.Color || compareCard.Value+1 != tmpCard.Value {
			return false
		}
		compareCard = tmpCard
	}

	*hv = combination.HandValue{
		Combination: combination.StraightFlush,
		Comparator: combination.StraightFlushComparator{
			HighestValue: h.cards[4].Value,
		},
	}
	return true
}

func (h hand) isFourOfAKind(hv *combination.HandValue) bool {
	valueCount := make(map[card.Value]int)
	for _, c := range h.cards {
		valueCount[c.Value]++
		if valueCount[c.Value] == 4 {
			*hv = combination.HandValue{
				Combination: combination.FourOfAKind,
				Comparator: combination.FourOfAKindComparator{
					Value: c.Value,
				},
			}
			return true
		}
	}
	return false
}

func (h hand) isFull(hv *combination.HandValue) bool {
	valueCount := make(map[card.Value]int)
	for _, c := range h.cards {
		valueCount[c.Value]++
	}
	//Four of a kind has to be checked before
	if len(valueCount) == 2 {
		var threeValue, pairValue card.Value
		for k, v := range valueCount {
			switch v {
			case 2:
				pairValue = k
			case 3:
				threeValue = k
			}
		}
		*hv = combination.HandValue{
			Combination: combination.Full,
			Comparator: combination.FullComparator{
				ThreeOfAKindValue: threeValue,
				PairValue:         pairValue,
			},
		}
		return true
	}
	return false
}

func (h hand) isFlush(hv *combination.HandValue) bool {
	compareCard := h.cards[0]
	for idx := 1; idx < 5; idx++ {
		tmpCard := h.cards[idx]
		if compareCard.Color != tmpCard.Color {
			return false
		}
		compareCard = tmpCard
	}

	*hv = combination.HandValue{
		Combination: combination.Flush,
		Comparator: combination.FlushComparator{
			HighestValue: h.cards[4].Value,
		},
	}
	return true
}

func (h hand) isStraight(hv *combination.HandValue) bool {
	compareCard := h.cards[0]
	for idx := 1; idx < 5; idx++ {
		tmpCard := h.cards[idx]
		if compareCard.Value+1 != tmpCard.Value {
			return false
		}
		compareCard = tmpCard
	}

	*hv = combination.HandValue{
		Combination: combination.Straight,
		Comparator: combination.StraightComparator{
			HighestValue: h.cards[4].Value,
		},
	}
	return true
}

func (h hand) isThreeOfAKind(hv *combination.HandValue) bool {
	valueCount := make(map[card.Value]int)
	for _, c := range h.cards {
		valueCount[c.Value]++
	}

	//Full has to be checked before
	for k, v := range valueCount {
		if v == 3 {
			*hv = combination.HandValue{
				Combination: combination.ThreeOfAKind,
				Comparator: combination.ThreeOfAKindComparator{
					Value: k,
				},
			}
			return true
		}
	}

	return false
}

func (h hand) isDoublePair(hv *combination.HandValue) bool {
	valueCount := make(map[card.Value]int)
	for _, c := range h.cards {
		valueCount[c.Value]++
	}

	possibleDoublePair := combination.DoublePairComparator{}
	pairCount := 0
	for k, v := range valueCount {
		switch v {
		case 1:
			possibleDoublePair.HighCardValue = k
		}
		if v == 2 {
			possibleDoublePair.PairValues[pairCount] = k
			pairCount++

			if pairCount == 2 {
				sortSliceDesc(possibleDoublePair.PairValues[:])
				*hv = combination.HandValue{
					Combination: combination.DoublePair,
					Comparator:  possibleDoublePair,
				}
				return true
			}
		}
	}

	return false
}

func (h hand) isPair(hv *combination.HandValue) bool {
	valueCount := make(map[card.Value]int)
	for _, c := range h.cards {
		valueCount[c.Value]++
	}

	if len(valueCount) != 4 {
		return false
	}

	var pairValue card.Value
	possibleHighCardValues := [3]card.Value{}
	singleCardCount := 0
	//Double pair has to be checked before
	for k, v := range valueCount {
		switch v {
		case 1:
			possibleHighCardValues[singleCardCount] = k
			singleCardCount++
		case 2:
			pairValue = k
		}
	}

	if pairValue.IsAValue() {
		sortSliceDesc(possibleHighCardValues[:])
		*hv = combination.HandValue{
			Combination: combination.Pair,
			Comparator: combination.PairComparator{
				PairValue:  pairValue,
				CardValues: possibleHighCardValues,
			},
		}
		return true
	}

	return false
}

func (h hand) isHighCard(hv *combination.HandValue) {
	values := h.Values()
	sortSliceDesc(values[:])
	*hv = combination.HandValue{
		Combination: combination.HighCard,
		Comparator:  combination.HighCardComparator{CardValues: values},
	}
}

func sortSliceDesc(values []card.Value) {
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
}
