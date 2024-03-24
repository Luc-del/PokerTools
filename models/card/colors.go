package card

//go:generate enumer -type Color -output colors.gen.go

// Color represents one of the four colors of a card game.
type Color int

const (
	Spade Color = iota + 1
	Club
	Diamond
	Heart
)
