package _021_04_07

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/4/8 12:05 上午
 * @description：
 * @modified By：
 * @version    ：$
 */
func numWays(n int) int {
	p, q, r := 0, 1, 1
	for i := 1; i <= n; i++ {
		r = (p + q) % (1e9 + 7)
		p = q
		q = r
	}
	return r
}
