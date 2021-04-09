package _021_04_09

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/4/9 9:51 上午
 * @description：
 * @modified By：
 * @version    ：$
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 记录当前节点
	curr := &ListNode{}
	// 哨兵节点
	head := curr
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	// 如果第一个数组没有遍历完
	for l1 != nil {
		curr.Next = l1
		curr = curr.Next
		l1 = l1.Next
	}
	// 如果第二个数组没有遍历完
	for l2 != nil {
		curr.Next = l2
		curr = curr.Next
		l2 = l2.Next
	}

	return head.Next
}
