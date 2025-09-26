// medium problem

package trees

func RightSideView(root *TreeNode) []int {
	res := []int{}
	level := [][]int{}

	var dfs func(node *TreeNode, lvl int)
	dfs = func(node *TreeNode, lvl int) {
		if node == nil {
			return
		}

		if len(level) == lvl {
			level = append(level, []int{})
		}

		level[lvl] = append(level[lvl], node.Val)

		dfs(node.Left, lvl+1)
		dfs(node.Right, lvl+1)
	}

	dfs(root, 0)

	for _, nodes := range level {
		res = append(res, nodes[len(nodes)-1])
	}

	return res
}
