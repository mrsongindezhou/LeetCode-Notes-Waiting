package _021_03_30

import "container/heap"

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/30 6:26 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, v := range nums {
		occurrences[v]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int            { return len(h) }
func (h IHeap) Less(i, j int) bool  { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}
