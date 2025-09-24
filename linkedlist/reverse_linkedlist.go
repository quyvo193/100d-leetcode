// easy problem
package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
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
