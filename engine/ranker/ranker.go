package ranker

import "pokerTools/models/hand"

type rankCheck func(h hand.Hand) (Strength, bool)

var rankCheckers = []rankCheck{
	isStraightFlush,
	isFourOfAKind,
	isFull,
	isFlush,
	isStraight,
	isThreeOfAKind,
	isDoublePair,
	isPair,
	isHighCard,
}

func ComputeStrength(h hand.Hand) Strength {
	for _, check := range rankCheckers {
		if s, ok := check(h); ok {
			return s
		}
	}

	// cannot happen as it should be at least a HighCard
	return ""
}
