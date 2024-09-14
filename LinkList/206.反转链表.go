package linklist

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