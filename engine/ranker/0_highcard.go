package ranker

import (
	"pokerTools/models/hand"
	"pokerTools/models/rank"
)

func highCard(h hand.Hand) Strength {
	return computePower(rank.HighCard, values(h)...)
}
