package _021_04_01

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/4/2 10:52 上午
 * @description：
 * @modified By：
 * @version    ：$
 */
func numIslands(grid [][]byte) int {
	landsNums := 0
	x, y := len(grid), len(grid[0])
	dfs := func(i, j int, grid [][]byte) {}
	dfs = func(i, j int, grid [][]byte) {
		if i < 0 || j < 0 || i >= x || j >= y || grid[i][j] == '0' {
			return
		}
		// 因为要进行上下左右递归，所以将已访问的置为0
		grid[i][j] = '0'
		dfs(i-1, j, grid)
		dfs(i+1, j, grid)
		dfs(i, j-1, grid)
		dfs(i, j+1, grid)
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			// 如果是1，则通过dfs找到不为0的为止
			if grid[i][j] == '1' {
				landsNums++
				dfs(i, j, grid)
			}
		}
	}
	return landsNums
}
