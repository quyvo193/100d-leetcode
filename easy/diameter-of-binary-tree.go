package easy

func DiameterOfBinaryTree(root *TreeNode) int {
	diameter := 0

	var maxDepth func(root *TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		maxLeft := maxDepth(node.Left)
		maxRight := maxDepth(node.Right)
		diameter = max(diameter, maxLeft+maxRight)

		return max(maxLeft, maxRight) + 1
	}

	maxDepth(root)

	return diameter
}
