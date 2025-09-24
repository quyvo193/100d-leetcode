// easy problem
package trees

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(MaxDepth(root.Left)+1, MaxDepth(root.Right)+1)
}
