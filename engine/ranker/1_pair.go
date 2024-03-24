package ranker

import (
	"pokerTools/models/card"
	"pokerTools/models/hand"
	"pokerTools/models/rank"
	"slices"
)

func isPair(h hand.Hand) (Strength, bool) {
	count := map[card.Value]int{
		h.Cards[0].Value: 1,
	}
	for i := 1; i < len(h.Cards); i++ {
		count[h.Cards[i].Value]++
	}

	if len(count) != 4 {
		return "", false
	}
	var (
		pair    card.Value
		singles = make([]card.Value, 0, 3)
	)
	for k, v := range count {
		switch v {
		case 2:
			pair = k
		case 1:
			singles = append(singles, k)
		default:
			return "", false
		}
	}

	slices.Sort(singles)
	return computePower(rank.Pair, pair, singles[2], singles[1], singles[0]), true
}
