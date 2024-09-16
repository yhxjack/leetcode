package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
	// 模拟的做法，使用循环获取k个链表的起始位置，进行反转，然后在拼接
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		// 保存下一个循环的头结点
		next := tail.Next
		// 新的头结点。需要复写，不能新建
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return dummy.Next
}
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for tail != prev {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	return tail, head
}