package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) (*ListNode, int) {
	var r *ListNode
	n := 0

	for head != nil {
		n++
		tmp := head.Next
		head.Next = r
		r = head
		head.Next = tmp
	}

	return r, n
}

func ReorderList(head *ListNode) {
	tmp := head
	rl, n := reverseList(tmp)

	for range n / 2 {
		tmp.Next = head
		tmp = tmp.Next
		tmp.Next = rl
		tmp = tmp.Next
	}
}
