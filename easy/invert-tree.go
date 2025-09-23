package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	var invert func(node *TreeNode)
	invert = func(node *TreeNode) {
		if node == nil {
			return
		}

		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp

		invert(node.Left)
		invert(node.Right)
	}

	invert(root)

	return root
}
