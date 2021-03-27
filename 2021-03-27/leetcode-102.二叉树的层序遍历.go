package _021_03_27

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS遍历
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)       // 存放返回结果
	queue := make([]*TreeNode, 1) // 声明一个队列
	queue[0] = root
	for len(queue) > 0 {
		n := len(queue) // 记录本层长度
		floor := []int{}
		for n > 0 {
			n--
			node := queue[0] // 取出队列第一个值
			queue = queue[1:]
			if node == nil {
				continue
			}
			floor = append(floor, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
		if len(floor) > 0 {
			res = append(res, floor)
		}
	}
	return res
}
