package _021_04_08

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/4/8 10:06 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 哨兵节点
	hair := &ListNode{Next: head}
	pre := hair
	for head != nil {
		tail := pre
		// k个一组进行反转
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = rev(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

// 头节点变尾节点，尾节点变为头节点
func rev(head, tail *ListNode) (*ListNode, *ListNode) {
	p := head
	prev := tail.Next
	for tail != prev {
		next := p.Next
		p.Next = prev
		prev = p
		p = next

	}
	return tail, head
}
