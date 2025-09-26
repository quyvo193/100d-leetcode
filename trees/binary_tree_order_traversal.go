package trees

func LevelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	visted := make(map[*TreeNode]bool)
	res := [][]int{}

	queue = append(queue, root)

	for len(queue) > 0 {
		arr := []int{}
		pop := queue[0]
		queue = queue[1:]

		visted[pop] = true

		if pop.Left != nil {
			arr = append(arr, pop.Left.Val)
		}

		if pop.Right != nil {
			arr = append(arr, pop.Right.Val)
		}

		res = append(res, arr)

		if !visted[pop] {
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}

			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
	}

	return res
}
