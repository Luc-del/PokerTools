package card

//go:generate enumer -type Value -output values.gen.go

// Value represents the value of a card, Ace being the top one.
type Value int

const (
	Two Value = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)
