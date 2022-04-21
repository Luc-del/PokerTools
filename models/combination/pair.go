package combination

import (
	"pokerTools/models/card"
	"sort"
)

type PairComparator struct {
	PairValue  card.Value
	CardValues [3]card.Value
}

func (p PairComparator) getValue() Value {
	sort.Slice(p.CardValues[:], func(i, j int) bool {
		return p.CardValues[i] > p.CardValues[j]
	})

	fakePairHand := append([]card.Value{p.PairValue}, p.CardValues[:]...)
	return computeHandCombinationValue(fakePairHand...)
}
