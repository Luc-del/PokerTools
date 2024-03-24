package ranker

import (
	"pokerTools/models/card"
	"pokerTools/models/hand"
	"pokerTools/models/rank"
)

func isFourOfAKind(h hand.Hand) (Strength, bool) {
	count := map[card.Value]int{
		h.Cards[0].Value: 1,
	}
	for i := 1; i < len(h.Cards); i++ {
		count[h.Cards[i].Value]++
	}

	if len(count) != 2 {
		return "", false
	}
	var fourOfAKindValue, fifthCardValue card.Value
	for k, v := range count {
		switch v {
		case 4:
			fourOfAKindValue = k
		case 1:
			fifthCardValue = k
		default:
			return "", false
		}
	}

	return computePower(rank.FourOfAKind, fourOfAKindValue, fifthCardValue), true
}
