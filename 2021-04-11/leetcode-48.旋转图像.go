package _021_04_11

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/4/11 4:56 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
// 这里建立坐标系只能竖直方向为x, 水平方向为y ，二维数组建立坐标系最好都这样建
// matrix[2][0]习惯上都会认为是坐标(2, 0)，否则会打乱数组的索引,需要体会一下
func rotate(matrix [][]int) {
	n := len(matrix)
	// 这里如果i的取值是(0,(n+1)/2)则j为(0,n/2) 两个值可以更换
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < n/2; j++ {
			// 顺时针更换坐标值
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] =
				matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]
		}
	}
}
