package combination

import "pokerTools/models/card"

type StraightComparator struct {
	HighestValue card.Value
}

func (s StraightComparator) getValue() Value {
	return Value(cardValueToHexString(s.HighestValue))
}
