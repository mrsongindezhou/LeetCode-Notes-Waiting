package _021_03_30

import "math"

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/30 11:32 下午
 * @description：兑换硬币
 * @modified By：
 * @version    ：$
 */

// https://leetcode-cn.com/problems/coin-change/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	maxNum := math.MaxInt32
	for i := 0; i < len(dp); i++ {
		dp[i] = maxNum
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			if i >= v && dp[i-v] != maxNum {
				dp[i] = min(dp[i-v]+1, dp[i])
			}
		}
	}
	if dp[amount] == maxNum {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
