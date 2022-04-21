package combination

import (
	"pokerTools/models/card"
)

type FlushComparator struct {
	HighestValue card.Value
}

func (f FlushComparator) getValue() Value {
	return Value(cardValueToHexString(f.HighestValue))
}
