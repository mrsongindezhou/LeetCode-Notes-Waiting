package _021_04_07

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/4/7 7:30 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n) // dp代表第n个位置最长递增序列数量
	dp[0] = 1
	maxCount := 1
	// 从第二个位置开始，对前面的数依次进行循环查找
	for i := 1; i < n; i++ {
		// 记录本次循环找到的最大值
		maxVal := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxVal = max(maxVal, dp[j])
			}
		}
		dp[i] = maxVal + 1
		maxCount = max(maxCount, dp[i])
	}
	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
