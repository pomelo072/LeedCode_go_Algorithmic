package Remove_Nth_Node_From_End_of_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	tempHead := &ListNode{0, head}
	slowPoint, fastPoint := tempHead, head
	for i := 0; i < n; i++ {
		fastPoint = fastPoint.Next
	}
	for ; fastPoint != nil; fastPoint = fastPoint.Next {
		slowPoint = slowPoint.Next

	}
	slowPoint.Next = slowPoint.Next.Next
	return tempHead.Next
}
