package Remove_Nth_Node_From_End_of_List

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

func Link2Array(head *ListNode) []int {
	var arr []int
	if head == nil {
		return []int{}
	}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return arr
}

func TestRemoveNthFromEnd(t *testing.T) {
	rawArr1, rawArr2, rawArr3 := []int{1, 2, 3, 4, 5}, []int{1}, []int{1, 2}
	head1, head2, head3 := Array2Link(rawArr1), Array2Link(rawArr2), Array2Link(rawArr3)
	t1, t2, t3 := RemoveNthFromEnd(head1, 2), RemoveNthFromEnd(head2, 1), RemoveNthFromEnd(head3, 1)
	arr1, arr2, arr3 := Link2Array(t1), Link2Array(t2), Link2Array(t3)
	if v1, v2, v3 := assert.Equal(t, []int{1, 2, 3, 5}, arr1), assert.Equal(t, []int{}, arr2), assert.Equal(t, []int{1}, arr3); v1 && v2 && v3 {
		t.Log(v1)
		t.Log(v2)
		t.Log(v3)
	} else {
		t.Log(v1)
		t.Log(v2)
		t.Log(v3)
		t.Error("Error")
	}
}
