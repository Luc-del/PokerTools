package rank

//go:generate enumer -type Rank -output rank.gen.go

// Rank represents a 5-cards rank defined by poker rules.
type Rank int

const (
	HighCard Rank = iota
	Pair
	DoublePair
	ThreeOfAKind
	Straight
	Flush
	Full
	FourOfAKind
	StraightFlush
)
