package linklist

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

// 两种做法，第一种迭代，第二种递归
// 迭代
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	temp := dummy
	for temp.Next != nil && temp.Next.Next != nil {
		// 记录第一个和第二个node进行拼接
		firstNode := temp.Next
		secondNode := temp.Next.Next

		// 重新进行链表的拼接
		temp.Next = secondNode
		// 两个节点互换位置
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode
		// temp节点需要循环后移，遍历和交换后续的链表
		temp = firstNode
	}
	return dummy.Next
}

// 递归
func swapPairsByRecursion(head *ListNode) *ListNode {
	// 递归退出条件
	if head == nil || head.Next == nil {
		return head
	}
	//head表示当前第一个节点，head.next表示交换后的心得节点
	newHead := head.Next
	// 需要将新节点之后的所有节点进行两两相交
	head.Next = swapPairsByRecursion(newHead.Next)
	// 将head拼接到newHead之后，完成第一个和第二个元素的反转拼接
	newHead.Next = head
	return newHead
}
