package deck

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func makeCombinationsIndexes(values []int, subsetSize int) <-chan []int {
	outputChan := make(chan []int, binomialTheorem(len(values), subsetSize))
	rSubsetMaker(outputChan, nil, values, subsetSize)
	close(outputChan)
	return outputChan
}

func rSubsetMaker(output chan<- []int, currentSubset, remainingValues []int, remainingSize int) {
	if remainingSize == 0 {
		output <- currentSubset
		return
	}

	for i := range remainingValues {
		rSubsetMaker(output, append(currentSubset, remainingValues[i]), remainingValues[i+1:], remainingSize-1)
	}
}

// binomialTheorem computes the number of combinations
func binomialTheorem(setSize, subsetSize int) int {
	numerator, denominator := 1, 1
	for k := 1; k <= setSize-subsetSize; k++ {
		numerator *= setSize - k + 1
		denominator *= k
	}

	return numerator / denominator
}

//This test was created to compute combinations stored in fiveOutOfSevenCombinationsIndexes
func Test_FiveOutOfSevenIndexesCombinations(t *testing.T) {
	t.Skip()

	values := []int{0, 1, 2, 3, 4, 5, 6}
	subsetSize := 5
	output := make([][]int, 0, binomialTheorem(len(values), subsetSize))
	combinationChan := makeCombinationsIndexes(values, subsetSize)
	for c := range combinationChan {
		output = append(output, c)
	}

	data, _ := json.MarshalIndent(output, "   ", "  ")
	require.NoError(t, ioutil.WriteFile("fiveOutOfSevenCombinationsIndexes.json", data, 0755))
}
