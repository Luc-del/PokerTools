package binomial_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"gonum.org/v1/gonum/stat/combin"
)

// This test was created to compute combinations stored in precomputed.go file.
func Test_FiveOutOfSevenIndexesCombinations(t *testing.T) {
	t.Skip()

	output := combin.Combinations(7, 5)

	data, err := json.Marshal(output)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile("fiveOutOfSevenCombinationsIndexes.json", data, 0755))
}
