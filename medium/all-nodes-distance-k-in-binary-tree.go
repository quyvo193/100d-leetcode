package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DistanceK(root *TreeNode, target *TreeNode, k int) []int {
	res := []int{}
	nodegraph := make(map[int][]*TreeNode)

	var dfs func(node, prev *TreeNode)
	dfs = func(node, prev *TreeNode) {
		if node == nil {
			return
		}

		nodegraph[node.Val] = []*TreeNode{node.Left, node.Right, prev}

		dfs(node.Left, node)
		dfs(node.Right, node)
	}

	dfs(root, nil)

	var dfs2 func(node, prev *TreeNode, count int)
	dfs2 = func(node, prev *TreeNode, count int) {
		if node == nil || count < 0 {
			return
		}

		if count == 0 {
			res = append(res, node.Val)
		}

		nodes := nodegraph[node.Val]

		for _, n := range nodes {
			if n != prev {
				dfs2(n, node, count-1)
			}
		}
	}

	dfs2(target, nil, k)

	return res
}
