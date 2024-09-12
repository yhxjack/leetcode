package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	input := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	fmt.Println(reverseList(input))
}

func reverseList(head *ListNode) *ListNode {
	// 迭代
	var dummy *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = dummy
		dummy = cur
		cur = next
	}
	return dummy
}