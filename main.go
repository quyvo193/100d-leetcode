package main

import (
	"fmt"
	"leetcode/medium"
)

func main() {
	head := &medium.ListNode{Val: 1, Next: &medium.ListNode{Val: 2, Next: &medium.ListNode{Val: 3, Next: &medium.ListNode{Val: 4, Next: &medium.ListNode{Val: 5, Next: &medium.ListNode{Val: 6, Next: &medium.ListNode{Val: 7, Next: &medium.ListNode{Val: 8, Next: nil}}}}}}}}
	medium.ReorderList(head)
	fmt.Println("r: ")
}
