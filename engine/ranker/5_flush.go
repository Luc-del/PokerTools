package ranker

import (
	"pokerTools/models/hand"
	"pokerTools/models/rank"
)

func isFlush(h hand.Hand) (Strength, bool) {
	for i := 0; i < len(h.Cards)-1; i++ {
		if h.Cards[i+1].Color != h.Cards[i].Color { // same color
			return "", false
		}
	}

	return computePower(rank.Flush, values(h)...), true
}
