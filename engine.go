package main

import (
	"fmt"
	"pokerTools/models/card"
	"pokerTools/models/hand"
)

//func ComputeWinRate(d deck.Deck) float64 {
//	h := d.ComputeBestHand()
//	return 0
//}
//
//func ComputePreFlopWinRate(c1, c2 [2]card.Card) float64 {
//
//}

func main() {
	me := [5]card.Card{
		{
			Value: card.Two,
			Color: card.Spade,
		},
		{
			Value: card.Five,
			Color: card.Club,
		},
		{
			Value: card.King,
			Color: card.Heart,
		},
		{
			Value: card.Jack,
			Color: card.Spade,
		},
		{
			Value: card.Seven,
			Color: card.Diamond,
		},
	}

	villain := [5]card.Card{
		{
			Value: card.Two,
			Color: card.Spade,
		},
		{
			Value: card.Five,
			Color: card.Club,
		},
		{
			Value: card.King,
			Color: card.Heart,
		},
		{
			Value: card.Queen,
			Color: card.Spade,
		},
		{
			Value: card.King,
			Color: card.Diamond,
		},
	}

	mine := hand.NewHand(me)
	op := hand.NewHand(villain)
	res := mine.Wins(op)
	fmt.Println(res)
}
