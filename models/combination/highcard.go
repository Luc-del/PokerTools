package combination

import (
	"pokerTools/models/card"
	"sort"
)

type HighCardComparator struct {
	CardValues [5]card.Value
}

func (h HighCardComparator) getValue() Value {
	sort.Slice(h.CardValues[:], func(i, j int) bool {
		return h.CardValues[i] > h.CardValues[j]
	})

	return computeHandCombinationValue(h.CardValues[:]...)
}
