package trees

func LevelOrder(root *TreeNode) [][]int {
	lvl := [][]int{}

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if len(lvl) <= depth {
			lvl = append(lvl, []int{node.Val})
		} else {
			lvl[depth] = append(lvl[depth], node.Val)
		}

		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	dfs(root, 0)

	return lvl
}
