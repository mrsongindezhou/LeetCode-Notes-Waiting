package _021_03_30

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/30 9:22 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

func rand7() int {
	return 0
}

func rand10() int {
	for {
		res := (rand7()-1)*7 + rand7()
		if res <= 40 {
			return res%10 + 1
		}
	}
}
