package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// var rlinkedlist *ListNode

	for head != nil {
		temp := head.Next
		// temp.Next = head
		// head = rlinkedlist
		fmt.Println("temp val: ", temp.Val)
		fmt.Println("head val: ", head.Val)
		head = head.Next
	}

	return head
}

func main() {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}

	reverseList(a)

}
