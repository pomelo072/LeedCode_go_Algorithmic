package Middle_of_the_Linked_List

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Array2Link(arr []int) *ListNode {
	head := new(ListNode)
	next := head
	n := len(arr)
	for i, v := range arr {
		if i == n-1 {
			next.Val = v
			next.Next = nil
		} else {
			next.Val = v
			next.Next = &ListNode{}
			next = next.Next
		}
	}
	return head
}

func TestMiddleNode(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 3, 4, 5, 6}
	head1 := Array2Link(arr1)
	head2 := Array2Link(arr2)
	middle1 := MiddleNode(head1)
	middle2 := MiddleNode(head2)
	if assert.Equal(t, 3, middle1.Val) {
		t.Log(middle1)
	} else {
		t.Log(middle1)
		t.Error("middle1 Error")
	}
	if assert.Equal(t, 4, middle2.Val) {
		t.Log(middle2)
	} else {
		t.Log(middle2)
		t.Error("middle2 Error")
	}
}

func BenchmarkMiddleNode(b *testing.B) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 3, 4, 5, 6}
	head1 := Array2Link(arr1)
	head2 := Array2Link(arr2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MiddleNode(head1)
		_ = MiddleNode(head2)
	}
	b.StopTimer()
}

func BenchmarkMiddleNode2(b *testing.B) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 3, 4, 5, 6}
	head1 := Array2Link(arr1)
	head2 := Array2Link(arr2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MiddleNode2(head1)
		_ = MiddleNode2(head2)
	}
	b.StopTimer()
}
