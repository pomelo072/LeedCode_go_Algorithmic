package Add_Two_Numbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	l3 := &ListNode{Val: 0}
	l4 := &ListNode{Val: 0}
	v1 := AddTwoNumbers(l1, l2)
	v2 := AddTwoNumbers(l3, l4)
	t.Log("v1: ")
	for v1 != nil {
		t.Log(v1.Val)
		v1 = v1.Next
	}
	t.Log(" Actully v1: 7,0,8")
	t.Log("v2: ")
	for v2 != nil {
		t.Log(v2.Val)
		v2 = v2.Next
	}
	t.Log(" Actully v2: 0")
}
