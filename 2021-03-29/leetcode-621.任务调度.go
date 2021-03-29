package _021_03_29

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/29 1:19 上午
 * @description：
 * @modified By：
 * @version    ：$
 */

// 参考题解：https://leetcode-cn.com/problems/task-scheduler/solution/tian-tong-si-lu-you-tu-kan-wan-jiu-dong-by-mei-jia/
// 参考题解：https://leetcode-cn.com/problems/task-scheduler/solution/ren-wu-diao-du-qi-golangda-bai-100yong-hu-de-jie-f/

func leastInterval(tasks []byte, n int) int {
	// 默认分为26个桶
	chars := [26]int{}
	// res记录返回结果
	res, count := 0, 0
	for i := 0; i < len(tasks); i++ {
		chars[tasks[i]-'A']++
		if res < chars[tasks[i]-'A'] {
			res = chars[tasks[i]-'A']
			count = 1
		} else if chars[tasks[i]-'A'] == res {
			count++
		}
	}
	// 如果调度时间为0或者任务类型小于冷却
	// res记录的是桶的深度，因为桶的最后一层有可能不满，所以此处是res - 1
	// n是冷却时间，因为是循环♻️执行，所以需要+1，如果这个时间小于任务
	//
	if n == 0 || ((res-1)*(n+1)+count < len(tasks)) {
		return len(tasks)
	}
	return (res-1)*(n+1) + count
}
