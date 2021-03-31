package _021_03_31

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/31 6:46 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, start, end int) {
	// i指针指向最后一个小于nums[end]
	i, j := start, start
	for j < end {
		if nums[j] < nums[end] && i != j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	nums[i], nums[end] = nums[end], nums[i]
	quickSort(nums, start, i-1)
	quickSort(nums, i+1, end)
}
