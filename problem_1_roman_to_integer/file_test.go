package problem_1_roman_to_integer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func romanToInt(input string) int {
	romanSymbols := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}

	currentResult := 0
	for i := 0; i < len(input); i++ {
		fmt.Printf("%v Equal to %v\n", string(input[i]), romanSymbols[string(input[i])])
		if i+1 < len(input) {
			if romanSymbols[string(input[i])] < romanSymbols[string(input[i+1])] {
				currentResult += romanSymbols[string(input[i+1])] - romanSymbols[string(input[i])]
				fmt.Printf("Current Result in (skip 1) %v: %v\n", i, currentResult)
				i++
				continue
			}
		}
		currentResult += romanSymbols[string(input[i])]
		fmt.Printf("Current Result in %v: %v\n", i, currentResult)

	}

	return currentResult
}

func TestCases(t *testing.T) {
	t.Run("Test case 1", func(t *testing.T) {
		result := romanToInt("III")
		assert.Equal(t, 3, result)
	})
	t.Run("Test case 2", func(t *testing.T) {
		result := romanToInt("LVIII")
		assert.Equal(t, 58, result)
	})
	t.Run("Test case 3", func(t *testing.T) {
		result := romanToInt("MCMXCIV")
		assert.Equal(t, 1994, result)
	})
}
