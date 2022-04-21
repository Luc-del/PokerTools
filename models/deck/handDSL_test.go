package deck

import "pokerTools/models/card"

type handDSL struct {
	hand
}

func newHandDSL() handDSL {
	return handDSL{
		hand: hand{
			cards: [5]card.Card{
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
			},
		},
	}
}

func (dsl handDSL) withColors(colors ...card.Color) handDSL {
	switch len(colors) {
	case 1:
		for k := range dsl.hand.cards {
			dsl.hand.cards[k].Color = colors[0]
		}
	case 5:
		for k := range dsl.hand.cards {
			dsl.hand.cards[k].Color = colors[k]
		}
	default:
		panic("dsl colors length must be 1 or 5")
	}

	return dsl
}

func (dsl handDSL) withValues(values ...card.Value) handDSL {
	switch len(values) {
	case 5:
		for k := range dsl.hand.cards {
			dsl.hand.cards[k].Value = values[k]
		}
	default:
		panic("dsl values length must 5")
	}

	return dsl
}

func (dsl handDSL) build() hand {
	return dsl.hand
}
