package problem_876_middle_of_linked_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func betterMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fmt.Println(slow)
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func middleNode(head *ListNode) *ListNode {
	count := countListNode(head)
	for i := 0; i < count; i++ {
		if i < count/2 {
			head = head.Next
			continue
		}
		break
	}

	return head
}

func countListNode(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}

	return count
}

func TestCase_1(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	tail := &ListNode{Val: slice[len(slice)-1]}
	midNode := &ListNode{}
	for i := len(slice) - 2; i >= 0; i-- {
		if i == len(slice)/2 {
			midNode.Val = slice[i]
			midNode.Next = tail
		}
		head := &ListNode{
			Val:  slice[i],
			Next: tail,
		}
		tail = head
	}

	assert.Equal(t, midNode, middleNode(tail))
}

func TestCase_2(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	tail := &ListNode{Val: slice[len(slice)-1]}
	midNode := &ListNode{}
	for i := len(slice) - 2; i >= 0; i-- {
		if i == len(slice)/2 {
			midNode.Val = slice[i]
			midNode.Next = tail
		}
		head := &ListNode{
			Val:  slice[i],
			Next: tail,
		}
		tail = head
	}

	assert.Equal(t, midNode, middleNode(tail))
}

func TestCase_3(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	tail := &ListNode{Val: slice[len(slice)-1]}
	midNode := &ListNode{}
	for i := len(slice) - 2; i >= 0; i-- {
		if i == len(slice)/2 {
			midNode.Val = slice[i]
			midNode.Next = tail
		}
		head := &ListNode{
			Val:  slice[i],
			Next: tail,
		}
		tail = head
	}

	assert.Equal(t, midNode, betterMiddleNode(tail))
}

func TestCase_4(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	tail := &ListNode{Val: slice[len(slice)-1]}
	midNode := &ListNode{}
	for i := len(slice) - 2; i >= 0; i-- {
		if i == len(slice)/2 {
			midNode.Val = slice[i]
			midNode.Next = tail
		}
		head := &ListNode{
			Val:  slice[i],
			Next: tail,
		}
		tail = head
	}

	assert.Equal(t, midNode, betterMiddleNode(tail))
}
