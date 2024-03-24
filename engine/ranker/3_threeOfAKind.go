package ranker

import (
	"pokerTools/models/card"
	"pokerTools/models/hand"
	"pokerTools/models/rank"
)

func isThreeOfAKind(h hand.Hand) (Strength, bool) {
	count := map[card.Value]int{
		h.Cards[0].Value: 1,
	}
	for i := 1; i < len(h.Cards); i++ {
		count[h.Cards[i].Value]++
	}

	if len(count) != 3 {
		return "", false
	}
	var threeOfAKindValue, fourthValue, fifthValue card.Value
	for k, v := range count {
		switch v {
		case 3:
			threeOfAKindValue = k
		case 1:
			fifthValue = k
			if k > fourthValue {
				fifthValue = fourthValue
				fourthValue = k
			}
		default:
			return "", false
		}
	}

	return computePower(rank.ThreeOfAKind, threeOfAKindValue, fourthValue, fifthValue), true
}
