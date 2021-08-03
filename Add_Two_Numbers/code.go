package Add_Two_Numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddNum(a, b, c int) (int, int) {
	num := a + b + c
	c1 := 0
	if num > 9 {
		c1 = 1
		num = num % 10
		return num, c1
	} else {
		return num, c1
	}
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	tail := new(ListNode)
	num, c := 0, 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		num, c = AddNum(n1, n2, c)
		if head == nil {
			head = &ListNode{Val: num}
			tail = head
		} else {
			tail.Next = &ListNode{Val: num}
			tail = tail.Next
		}
	}
	if c > 0 {
		tail.Next = &ListNode{Val: c}
	}
	return
}
