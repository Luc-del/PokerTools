package combination

import (
	"fmt"
	"pokerTools/models/card"
)

type HandValue struct {
	Combination
	Comparator
}

type Comparator interface {
	getValue() Value
}

// Value represent the hexadecimal string highestValue of the concatenation of a given card
// Comparing Value will be on most significant bit
type Value string

func computeHandCombinationValue(cardValues ...card.Value) Value {
	hexStringConcat := ""
	for _, value := range cardValues {
		//hex conversion
		hexStringConcat += cardValueToHexString(value)
	}

	return Value(hexStringConcat)
}

func cardValueToHexString(v card.Value) string {
	return fmt.Sprintf("%x", int(v))
}
