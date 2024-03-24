package binomial

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

//These tests were created to compute combinations stored in precomputed.go file

func Test_FiveOutOfSevenIndexesCombinations(t *testing.T) {
	t.Skip()

	values := []int{0, 1, 2, 3, 4, 5, 6}
	subsetSize := 5
	output := make([][]int, 0, Coefficient(len(values), subsetSize))
	combinationChan := MakeCombinationsIndexes(values, subsetSize)
	for c := range combinationChan {
		output = append(output, c)
	}

	data, _ := json.MarshalIndent(output, "   ", "  ")
	require.NoError(t, ioutil.WriteFile("fiveOutOfSevenCombinationsIndexes.json", data, 0755))
}

func Test_FiveOutOfFortyEightIndexesCombinations(t *testing.T) {
	t.Skip()

	var values []int
	for i := 0; i < 48; i++ {
		values = append(values, i)
	}
	subsetSize := 5
	output := make([][]int, 0, Coefficient(len(values), subsetSize))
	combinationChan := MakeCombinationsIndexes(values, subsetSize)
	for c := range combinationChan {
		output = append(output, c)
	}

	data, _ := json.MarshalIndent(output, "", "")
	require.NoError(t, os.WriteFile("fiveOutOfFortyEightCombinationsIndexes.json", data, 0755))
}
