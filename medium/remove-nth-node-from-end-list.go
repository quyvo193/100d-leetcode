package medium

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	first, d := head, head

	for i := 0; i < n; i++ {
		first = first.Next
	}

	if first == nil {
		return head.Next
	}

	for first.Next != nil {
		d = d.Next
		first = first.Next
	}

	if d.Next == nil {
		head.Next = nil
		return head
	}

	d.Next = d.Next.Next

	return head
}
