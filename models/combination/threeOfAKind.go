package combination

import "pokerTools/models/card"

type ThreeOfAKindComparator struct {
	Value card.Value
}

func (t ThreeOfAKindComparator) getValue() Value {
	return Value(cardValueToHexString(t.Value))
}
