package combination

import "pokerTools/models/card"

type StraightFlushComparator struct {
	HighestValue card.Value
}

func (s StraightFlushComparator) getValue() Value {
	return Value(cardValueToHexString(s.HighestValue))
}
