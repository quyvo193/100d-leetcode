package easy

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	res := true

	var isSame func(p *TreeNode, q *TreeNode)
	isSame = func(p *TreeNode, q *TreeNode) {
		if (q == nil && p == nil) || !res {
			return
		}

		if (p == nil && q != nil) || (p != nil && q == nil) || (p.Val != q.Val) {
			res = false
			return
		}

		isSame(p.Left, q.Left)
		isSame(p.Right, q.Right)
	}

	isSame(p, q)

	return res
}
