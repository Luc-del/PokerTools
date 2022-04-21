package combination

import "pokerTools/models/card"

type FullComparator struct {
	ThreeOfAKindValue card.Value
	PairValue         card.Value
}

func (f FullComparator) getValue() Value {
	return computeHandCombinationValue(f.ThreeOfAKindValue, f.PairValue)
}
