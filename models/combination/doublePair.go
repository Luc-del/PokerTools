package combination

import (
	"pokerTools/models/card"
	"sort"
)

type DoublePairComparator struct {
	PairValues    [2]card.Value
	HighCardValue card.Value
}

func (dp DoublePairComparator) getValue() Value {
	sort.Slice(dp.PairValues[:], func(i, j int) bool {
		return dp.PairValues[i] > dp.PairValues[j]
	})

	return computeHandCombinationValue(dp.PairValues[0], dp.PairValues[1], dp.HighCardValue)
}
