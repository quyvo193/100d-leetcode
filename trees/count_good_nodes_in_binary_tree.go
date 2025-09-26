// medium problem
package trees

func GoodNodes(root *TreeNode) int {
	count := 0

	var dfs func(node *TreeNode, min int)
	dfs = func(node *TreeNode, min int) {
		if node == nil {
			return
		}

		if node.Val >= min {
			count++
			min = node.Val
		}

		dfs(node.Left, min)
		dfs(node.Right, min)
	}

	dfs(root, root.Val)

	return count
}
