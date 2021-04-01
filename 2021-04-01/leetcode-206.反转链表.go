package _021_04_01

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/4/1 7:12 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归实现
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// last节点是链表对最后一个节点，通过每一层递归向上传递
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// 迭代实现
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}
