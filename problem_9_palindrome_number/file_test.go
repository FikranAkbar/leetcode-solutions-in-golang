package problem_9_palindrome_number

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func isPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)
	var reverseStr string
	for j := len(str) - 1; j >= 0; j = j - 1 {
		reverseStr = reverseStr + string(str[j])
	}

	return str == reverseStr
}

func TestPalindromeTest(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, true, isPalindromeNumber(121))
	})
	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, false, isPalindromeNumber(-121))
	})
	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, false, isPalindromeNumber(10))
	})
	t.Run("Test 4", func(t *testing.T) {
		assert.Equal(t, false, isPalindromeNumber(1234231))
	})
}
