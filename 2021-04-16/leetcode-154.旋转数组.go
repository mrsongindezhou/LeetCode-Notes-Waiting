package _021_04_16

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/4/16 10:23 上午
 * @description：
 * @modified By：
 * @version    ：$
 */
func minArray(numbers []int) int {
	low, high := 0, len(numbers)-1
	for low < high {
		// 取中间值
		pivot := low + (high-low)/2
		// 中间节点小于右侧节点说明，右侧区间是符合要求的，拐点在左侧，因为我们知道右侧旋转数组是小于左侧的
		if numbers[pivot] < numbers[high] {
			//pivot是已知符合条件的所以不用再次列为目标区间，因为判断条件是low < high所以此处是 hign = pivot
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			// 正常序列是自左向右增大，如果在中间值右侧出现了拐点，则舍弃左侧区间
			low = pivot + 1 // 因为pivot是已知结果，所以不用列为待查询区间
		} else {
			high--
		}
	}
	return numbers[low]
}
