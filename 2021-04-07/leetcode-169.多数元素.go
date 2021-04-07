package _021_04_07

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/4/7 10:35 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func majorityElement(nums []int) int {
	numsMap := map[int]int{}
	for _, v := range nums {
		numsMap[v]++
		if numsMap[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}
