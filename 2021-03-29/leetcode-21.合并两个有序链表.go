package _021_03_29

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/29 5:45 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{} // 哨兵节点，方便返回数据
	node := head        // 记录当前的节点
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			node.Next = l2
			l2 = l2.Next
		} else {
			node.Next = l1
			l1 = l1.Next
		}
		node = node.Next
	}
	// 如果l1链表不为空，将l1数据都合并到结果集中
	if l1 != nil {
		node.Next = l1
		node = node.Next
		l1 = l1.Next
	}
	// 如果l2链表不为空，将l2数据都合并到结果集中
	if l2 != nil {
		node.Next = l2
		node = node.Next
		l2 = l2.Next
	}
	return head.Next
}
