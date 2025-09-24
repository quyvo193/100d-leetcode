// easy problem
package trees

func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Val == subRoot.Val && isSameTree(root, subRoot) {
		return true
	}

	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
