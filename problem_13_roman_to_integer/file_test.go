package problem_13_roman_to_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func romanToInt(input string) int {
	romanSymbols := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}

	currentResult := 0
	for i := 0; i < len(input); i++ {
		if i+1 < len(input) {
			if romanSymbols[string(input[i])] < romanSymbols[string(input[i+1])] {
				currentResult += romanSymbols[string(input[i+1])] - romanSymbols[string(input[i])]
				i++
				continue
			}
		}
		currentResult += romanSymbols[string(input[i])]
	}

	return currentResult
}

func romanToIntReverseVersion(input string) int {
	romanSymbols := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}

	currentResult := romanSymbols[string(input[len(input)-1])]
	for i := len(input) - 2; i >= 0; i-- {
		if romanSymbols[string(input[i])] < romanSymbols[string(input[i+1])] {
			currentResult -= romanSymbols[string(input[i])]
		} else {
			currentResult += romanSymbols[string(input[i])]
		}
	}

	return currentResult
}

func TestNormalVersion(t *testing.T) {
	t.Run("Test case 1", func(t *testing.T) {
		result := romanToIntReverseVersion("III")
		assert.Equal(t, 3, result)
	})
	t.Run("Test case 2", func(t *testing.T) {
		result := romanToIntReverseVersion("LVIII")
		assert.Equal(t, 58, result)
	})
	t.Run("Test case 3", func(t *testing.T) {
		result := romanToIntReverseVersion("MCMXCIV")
		assert.Equal(t, 1994, result)
	})
}

func TestReverseVersion(t *testing.T) {
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
