// medium problem
package trees

import "math"

func IsValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, min, max int) bool
	dfs = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}

		if node.Val >= max || node.Val <= min {
			return false
		}

		return dfs(node.Left, min, node.Val) && dfs(node.Right, node.Val, max)
	}

	return dfs(root, math.MinInt, math.MaxInt)
}
