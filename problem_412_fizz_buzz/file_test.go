package problem_412_fizz_buzz

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		value := strconv.Itoa(i)
		if i%3 == 0 && i%5 == 0 {
			value = "FizzBuzz"
		} else if i%3 == 0 {
			value = "Fizz"
		} else if i%5 == 0 {
			value = "Buzz"
		}
		result[i-1] = value
	}

	return result
}

func TestCase_1(t *testing.T) {
	assert.Equal(t,
		[]string{"1", "2", "Fizz"},
		fizzBuzz(3),
	)
}

func TestCase_2(t *testing.T) {
	assert.Equal(t,
		[]string{"1", "2", "Fizz", "4", "Buzz"},
		fizzBuzz(5),
	)
}

func TestCase_3(t *testing.T) {
	assert.Equal(t,
		[]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		fizzBuzz(15),
	)
}
