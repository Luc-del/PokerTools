package combination

import "pokerTools/models/card"

type FourOfAKindComparator struct {
	Value card.Value
}

func (f FourOfAKindComparator) getValue() Value {
	return Value(cardValueToHexString(f.Value))
}
