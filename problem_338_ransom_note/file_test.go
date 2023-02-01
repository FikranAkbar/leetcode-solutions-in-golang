package problem_338_ransom_note

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteMap := convertStringToMap(ransomNote)
	magazineMap := convertStringToMap(magazine)

	if len(ransomNote) > len(magazine) {
		return false
	}

	for i, ransomNoteLetter := range ransomNote {
		for j, magazineLetter := range magazine {
			if ransomNoteLetter == magazineLetter && !magazineMap[j] {
				fmt.Printf("Comparing %v to %v is true, make ransomNoteMap index %v to true\n", string(ransomNoteLetter), string(magazineLetter), i)
				ransomNoteMap[i] = true
				magazineMap[j] = true
				break
			}
		}

		if !ransomNoteMap[i] {
			return false
		}
	}

	return true
}

func convertStringToMap(input string) map[int]bool {
	currentMap := map[int]bool{}
	for index, _ := range input {
		currentMap[index] = false
	}

	return currentMap
}

func TestCase_1(t *testing.T) {
	assert.Equal(t, false, canConstruct("a", "b"))
}

func TestCase_2(t *testing.T) {
	assert.Equal(t, false, canConstruct("aa", "ab"))
}

func TestCase_3(t *testing.T) {
	assert.Equal(t, true, canConstruct("aa", "aab"))
}
