// medium problem
package linkedlist

func HasCycle(head *ListNode) bool {
	addrMap := make(map[*ListNode]bool)

	for head != nil {
		addrMap[head] = true

		if addrMap[head.Next] {
			return true
		}

		head = head.Next
	}

	return false
}
