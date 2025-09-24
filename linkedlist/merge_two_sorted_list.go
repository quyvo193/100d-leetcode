// medium problem
package linkedlist

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergeList := &ListNode{}
	tmp := mergeList

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp.Next = list1
			list1 = list1.Next
		} else {
			tmp.Next = list2
			list2 = list2.Next
		}
		tmp = tmp.Next
	}

	if list1 == nil {
		tmp.Next = list2
	} else {
		tmp.Next = list1
	}

	return mergeList.Next
}
