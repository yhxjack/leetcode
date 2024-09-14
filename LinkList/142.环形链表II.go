package linklist

// 环形链表II
/* 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
不允许修改 链表。
*/

// 可以使用两种做法，第一种使用hashmap，判断有没有环，第二种使用快慢指针

// hashmap
func detectCycleByMap(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	st := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := st[head]; ok {
			return head
		}
		st[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func detectCycleByPoint(head *ListNode) *ListNode {
	// 使用快慢指针
	if head == nil {
		return nil
	}
fast, slow := head, head
	for fast != nil {
		slow = slow.Next
		// 若无环则返回nil
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		// 若fast==slow则说明有环
		if fast == slow {
			temp := head
			for temp != slow {
				temp = temp.Next
				slow = slow.Next
			}
			return temp
		}
	}
	return nil
}
