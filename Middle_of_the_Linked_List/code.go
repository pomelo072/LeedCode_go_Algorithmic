package Middle_of_the_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	n := 0
	count := head
	for count != nil {
		n++
		count = count.Next
	}
	middle := n / 2
	for i := 0; i < middle; i++ {
		head = head.Next
	}
	return head
}
func MiddleNode2(head *ListNode) *ListNode {
	slowPoint, fastPoint := head, head
	for fastPoint != nil && fastPoint.Next != nil {
		slowPoint = slowPoint.Next
		fastPoint = fastPoint.Next.Next
	}
	return slowPoint
}
