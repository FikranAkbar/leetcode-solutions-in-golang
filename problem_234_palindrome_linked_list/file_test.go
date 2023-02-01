package problem_234_palindrome_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var slice []int
	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}

	return isSlicePalindrome(slice)
}

func isSlicePalindrome(slice []int) bool {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		if slice[i] != slice[j] {
			return false
		}
	}

	return true
}

func TestCase_1(t *testing.T) {
	slice := []int{1, 2, 2, 1}
	tail := &ListNode{Val: slice[len(slice)-1]}
	for i := len(slice) - 2; i >= 0; i-- {
		head := &ListNode{
			Val:  slice[i],
			Next: tail,
		}
		tail = head
	}

	assert.Equal(t, true, isPalindrome(tail))
}

func TestCase_2(t *testing.T) {
	slice := []int{1, 2}
	tail := &ListNode{Val: slice[len(slice)-1]}
	for i := len(slice) - 2; i >= 0; i-- {
		head := &ListNode{
			Val:  slice[i],
			Next: tail,
		}
		tail = head
	}

	assert.Equal(t, false, isPalindrome(tail))
}
