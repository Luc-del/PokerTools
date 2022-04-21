package combination

//go:generate enumer -type Combination -output combinations.gen.go

type Combination int

const (
	HighCard Combination = iota
	Pair
	DoublePair
	ThreeOfAKind
	Straight
	Flush
	Full
	FourOfAKind
	StraightFlush
)
