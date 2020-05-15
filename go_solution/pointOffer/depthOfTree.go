package pointOffer

// BFS，用队列实现
func maxDepth1(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	queue := []*TreeNode{root}
	var curr *TreeNode
	for len(queue) != 0 {
		depth ++
		n := len(queue)
		for i:= 0; i < n; i++ {
			curr, queue = queue[0], queue[1:]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}
	return depth
}

// DFS 递归方式
func maxDepth2(root *TreeNode) int{
	return depth(root)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(depth(root.Left), depth(root.Right))
}

func max(a, b int )int {
	if a> b {
		return a
	}
	return b
}

