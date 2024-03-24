package ranker

import (
	"pokerTools/models/hand"
	"pokerTools/models/rank"
)

func isHighCard(h hand.Hand) (Strength, bool) {
	return computePower(rank.HighCard, values(h)...), true
}
