package binaryTreeLevelOrderTraversalII

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	queue := []*TreeNode{root}
	curr := root
	for len(queue) != 0 {
		tmp := make([]int, 0)
		level := make([]*TreeNode, 0)
		for len(queue) != 0 {
			curr, queue = queue[0], queue[1:]
			tmp = append(tmp, curr.Val)
			if curr.Left != nil {
				level = append(level, curr.Left)
			}
			if curr.Right != nil {
				level = append(level, curr.Right)
			}
		}
		queue = level
		ret = append([][]int{tmp}, ret...)
	}
	return ret
}
