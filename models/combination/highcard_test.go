package combination

import (
	"github.com/stretchr/testify/assert"
	"pokerTools/models/card"
	"testing"
)

func Test_highCard_getValue(t *testing.T) {
	tests := []struct {
		name  string
		cards [5]card.Value
		want  Value
	}{
		{
			name: "given sorted CardValues - top to low",
			cards: [5]card.Value{
				card.Ace,
				card.Queen,
				card.Nine,
				card.Eight,
				card.Three,
			},
			want: "ec983",
		},
		{
			name: "given sorted CardValues - low to top",
			cards: [5]card.Value{
				card.Three,
				card.Eight,
				card.Nine,
				card.Queen,
				card.Ace,
			},
			want: "ec983",
		},
		{
			name: "given unsorted CardValues",
			cards: [5]card.Value{
				card.Eight,
				card.Queen,
				card.Nine,
				card.Ace,
				card.Three,
			},
			want: "ec983",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			h := HighCardComparator{
				CardValues: testCase.cards,
			}
			assert.Equal(t, testCase.want, h.getValue())
		})
	}
}
