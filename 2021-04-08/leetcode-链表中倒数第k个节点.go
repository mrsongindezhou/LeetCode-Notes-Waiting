package _021_04_08

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/4/8 10:57 上午
 * @description：
 * @modified By：
 * @version    ：$
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	f, l := head, head
	for i := 0; i < k && f != nil; i++ {
		f = f.Next
	}
	for f != nil {
		f = f.Next
		l = l.Next
	}
	return l
}
