// Code generated by "enumer -type Rank -output rank.gen.go"; DO NOT EDIT.

package rank

import (
	"fmt"
	"strings"
)

const _RankName = "HighCardPairDoublePairThreeOfAKindStraightFlushFullFourOfAKindStraightFlush"

var _RankIndex = [...]uint8{0, 8, 12, 22, 34, 42, 47, 51, 62, 75}

const _RankLowerName = "highcardpairdoublepairthreeofakindstraightflushfullfourofakindstraightflush"

func (i Rank) String() string {
	if i < 0 || i >= Rank(len(_RankIndex)-1) {
		return fmt.Sprintf("Rank(%d)", i)
	}
	return _RankName[_RankIndex[i]:_RankIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _RankNoOp() {
	var x [1]struct{}
	_ = x[HighCard-(0)]
	_ = x[Pair-(1)]
	_ = x[DoublePair-(2)]
	_ = x[ThreeOfAKind-(3)]
	_ = x[Straight-(4)]
	_ = x[Flush-(5)]
	_ = x[Full-(6)]
	_ = x[FourOfAKind-(7)]
	_ = x[StraightFlush-(8)]
}

var _RankValues = []Rank{HighCard, Pair, DoublePair, ThreeOfAKind, Straight, Flush, Full, FourOfAKind, StraightFlush}

var _RankNameToValueMap = map[string]Rank{
	_RankName[0:8]:        HighCard,
	_RankLowerName[0:8]:   HighCard,
	_RankName[8:12]:       Pair,
	_RankLowerName[8:12]:  Pair,
	_RankName[12:22]:      DoublePair,
	_RankLowerName[12:22]: DoublePair,
	_RankName[22:34]:      ThreeOfAKind,
	_RankLowerName[22:34]: ThreeOfAKind,
	_RankName[34:42]:      Straight,
	_RankLowerName[34:42]: Straight,
	_RankName[42:47]:      Flush,
	_RankLowerName[42:47]: Flush,
	_RankName[47:51]:      Full,
	_RankLowerName[47:51]: Full,
	_RankName[51:62]:      FourOfAKind,
	_RankLowerName[51:62]: FourOfAKind,
	_RankName[62:75]:      StraightFlush,
	_RankLowerName[62:75]: StraightFlush,
}

var _RankNames = []string{
	_RankName[0:8],
	_RankName[8:12],
	_RankName[12:22],
	_RankName[22:34],
	_RankName[34:42],
	_RankName[42:47],
	_RankName[47:51],
	_RankName[51:62],
	_RankName[62:75],
}

// RankString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func RankString(s string) (Rank, error) {
	if val, ok := _RankNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _RankNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Rank values", s)
}

// RankValues returns all values of the enum
func RankValues() []Rank {
	return _RankValues
}

// RankStrings returns a slice of all String values of the enum
func RankStrings() []string {
	strs := make([]string, len(_RankNames))
	copy(strs, _RankNames)
	return strs
}

// IsARank returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Rank) IsARank() bool {
	for _, v := range _RankValues {
		if i == v {
			return true
		}
	}
	return false
}
