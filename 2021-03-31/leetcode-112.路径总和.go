package _021_03_31

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/31 7:50 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归实现
func hasPathSum(root *TreeNode, targetSum int) bool {
	// 递归结束条件
	if root == nil {
		return false
	}
	targetSum = targetSum - root.Val
	// 递归结束条件，符合
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
