package river

import (
	"pokerTools/binomial"
	"pokerTools/models/card"
	"pokerTools/models/hand"
)

type River struct {
	Cards [7]card.Card
}

func (d River) ComputeBestHand() hand.Hand {
	hands := d.makeHands()
	bestHand := hands[0]
	for idx := 1; idx < len(hands); idx++ {
		currentHand := hands[idx]
		if currentHand.Wins(bestHand) {
			bestHand = currentHand
		}
	}

	return bestHand
}

func (d River) makeHands() (hands [21]hand.Hand) {
	for handIndex, indexes := range binomial.FiveOutOfSevenCombinationsIndexes {
		hands[handIndex] = d.makeSingleHandFromIndexes(indexes)
	}

	return hands
}

func (d River) makeSingleHandFromIndexes(cardIndexes [5]int) hand.Hand {
	cards := [5]card.Card{}
	for k, idx := range cardIndexes {
		cards[k] = d.Cards[idx]
	}

	return hand.NewHand(cards)
}
