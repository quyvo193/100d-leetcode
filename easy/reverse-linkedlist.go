package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var rlinkedlist *ListNode

	for head != nil {
		temp := head.Next
		head.Next = rlinkedlist
		rlinkedlist = head
		head = temp
	}

	return head
}
