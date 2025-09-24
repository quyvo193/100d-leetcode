// easy problem
package linkedlist

func ReorderList(head *ListNode) {
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// reverse list
	second := slow.Next
	slow.Next = nil
	var rev *ListNode

	for second != nil {
		tmp := second.Next
		second.Next = rev
		rev = second
		second = tmp
	}

	// merge list

	first, second := head, rev

	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}
}
