package _021_04_02

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/4/2 3:41 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
type MyQueue struct {
	Stack1 []int
	Stack2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		Stack1: make([]int, 0),
		Stack2: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Stack1 = append(this.Stack1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	s2Len := len(this.Stack2)
	if s2Len > 0 {
		r := this.Stack2[s2Len-1]
		this.Stack2 = this.Stack2[:s2Len-1]
		return r
	} else {
		n := len(this.Stack1)
		if n > 0 {
			for i := n - 1; i > 0; i-- {
				this.Stack2 = append(this.Stack2, this.Stack1[i])
			}
			r := this.Stack1[0]
			this.Stack1 = this.Stack1[:0]
			return r
		}
	}
	return 0
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.Stack2) > 0 {
		return this.Stack2[len(this.Stack2)-1]
	}
	n := len(this.Stack1)
	if n > 0 {
		for i := n - 1; i >= 0; i-- {
			this.Stack2 = append(this.Stack2, this.Stack1[i])
		}
		this.Stack1 = this.Stack1[:0]
		return this.Stack2[len(this.Stack2)-1]
	}
	return 0
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.Stack1) == 0 && len(this.Stack2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
