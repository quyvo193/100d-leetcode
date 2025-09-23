package easy

func IsBalanced(root *TreeNode) bool {
	res := true

	var maxDepth func(node *TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l, r := maxDepth(node.Left), maxDepth(node.Right)

		if l-r > 1 || r-l > 1 {
			res = false
		}

		return max(l, r) + 1
	}

	maxDepth(root)

	return res
}
