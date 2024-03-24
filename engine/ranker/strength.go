package ranker

import (
	"fmt"
	"pokerTools/models/card"
	"pokerTools/models/hand"
	"pokerTools/models/rank"
	"strconv"
)

// Strength represents the strength of a given hand.Hand.
// It is represented by a hexadecimal string, allowing to handle card.Value >= 10.
// Comparing Strength will be done on the most significant bit.
type Strength string

func computePower(r rank.Rank, cardValues ...card.Value) Strength {
	hexStringConcat := strconv.Itoa(int(r))
	for _, value := range cardValues {
		hexStringConcat += cardValueToHexString(value) // hex conversion
	}

	return Strength(hexStringConcat)
}

func cardValueToHexString(v card.Value) string {
	return fmt.Sprintf("%x", int(v))
}

func values(h hand.Hand) []card.Value {
	v := make([]card.Value, 5)
	for i := range h.Cards {
		v[i] = h.Cards[i].Value
	}
	return v
}
