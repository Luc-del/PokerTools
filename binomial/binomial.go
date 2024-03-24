package binomial

func MakeCombinationsIndexes(values []int, subsetSize int) <-chan []int {
	outputChan := make(chan []int, Coefficient(len(values), subsetSize))
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

// Coefficient computes the number of combinations
// see https://en.wikipedia.org/wiki/Binomial_coefficient
func Coefficient(n, p int) int {
	res := 1.
	for k := 1; k <= n-max(p, n-p); k++ {
		res *= float64(n+1)/float64(k) - 1
	}

	return int(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
