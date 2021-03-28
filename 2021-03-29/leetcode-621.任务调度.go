package _021_03_29

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/29 1:19 上午
 * @description：
 * @modified By：
 * @version    ：$
 */
func leastInterval(tasks []byte, n int) int {
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
	if n == 0 || ((res-1)*(n+1)+count < len(tasks)) {
		return len(tasks)
	}
	return (res-1)*(n+1) + count
}
