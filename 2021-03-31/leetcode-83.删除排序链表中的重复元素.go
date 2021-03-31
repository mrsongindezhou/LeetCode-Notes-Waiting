package _021_03_31

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/31 4:36 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{}
	var pre *ListNode = head
	newHead.Next = pre
	head = head.Next
	for head != nil {
		if head.Val == pre.Val {
			pre.Next = head.Next
		} else {
			pre = head
		}
		head = head.Next
	}
	return newHead.Next
}

// 不借助辅助节点的写法
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	// 因为要判断当前节点和下一个节点，所以此处使用cur.Next作为判断条件
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
