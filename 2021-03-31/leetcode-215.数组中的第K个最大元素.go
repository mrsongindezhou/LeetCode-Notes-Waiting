package _021_03_31

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/30 6:15 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	// 要找第K大，则把比K大的都删掉，则堆顶就是我们要找的值，注意此处 len(nums) - k + 1
	// 例如：1、2、3 我们要找第2大，则3 - 2 + 1
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		// 这里并没有左真的删除，而是将堆顶与堆尾元素做调换，使堆做调整，达到效果
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		// 对0对应的堆节点进行递归调整
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

// 构建大顶堆
func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
}

// 递归调整
func maxHeapify(nums []int, i, heapSize int) {
	// 左子树，右子树
	l, r, large := 2*i+1, 2*i+2, i
	// 如果左子树大于父节点，则记录左子树的索引
	if l < heapSize && nums[large] < nums[l] {
		large = l
	}
	// 如果右子树大于父节点，则记录右子树的索引
	if r < heapSize && nums[large] < nums[r] {
		large = r
	}
	// 两个值不想等，证明发生了调整，则递归调整下层结构
	if i != large {
		nums[i], nums[large] = nums[large], nums[i]
		maxHeapify(nums, large, heapSize)
	}
}
