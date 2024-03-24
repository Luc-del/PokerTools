package ranker

import (
	"pokerTools/models/card"
	"pokerTools/models/hand"
	"pokerTools/models/rank"
)

func isFull(h hand.Hand) (Strength, bool) {
	count := map[card.Value]int{
		h.Cards[0].Value: 1,
	}
	for i := 1; i < len(h.Cards); i++ {
		count[h.Cards[i].Value]++
	}

	if len(count) != 2 {
		return "", false
	}
	var threeOfAKindValue, pairValue card.Value
	for k, v := range count {
		switch v {
		case 3:
			threeOfAKindValue = k
		case 2:
			pairValue = k
		default:
			return "", false
		}
	}

	return computePower(rank.Full, threeOfAKindValue, pairValue), true
}
