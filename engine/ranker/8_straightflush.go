package ranker

import (
	"pokerTools/models/hand"
	"pokerTools/models/rank"
)

func isStraightFlush(h hand.Hand) (Strength, bool) {
	for i := 0; i < len(h.Cards)-1; i++ {
		if h.Cards[i+1].Color != h.Cards[i].Color || // same color
			h.Cards[i+1].Value != h.Cards[i].Value-1 { // incremental value
			return "", false
		}
	}
	return computePower(rank.StraightFlush, h.Cards[0].Value), true
}
