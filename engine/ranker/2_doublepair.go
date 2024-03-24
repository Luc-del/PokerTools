package ranker

import (
	"pokerTools/models/card"
	"pokerTools/models/hand"
	"pokerTools/models/rank"
)

func isDoublePair(h hand.Hand) (Strength, bool) {
	count := map[card.Value]int{
		h.Cards[0].Value: 1,
	}
	for i := 1; i < len(h.Cards); i++ {
		count[h.Cards[i].Value]++
	}

	if len(count) != 3 {
		return "", false
	}
	var topPair, bottomPair, fifthValue card.Value
	for k, v := range count {
		switch v {
		case 2:
			bottomPair = k
			if k > topPair {
				bottomPair = topPair
				topPair = k
			}
		case 1:
			fifthValue = k
		default:
			return "", false
		}
	}

	return computePower(rank.DoublePair, topPair, bottomPair, fifthValue), true
}
