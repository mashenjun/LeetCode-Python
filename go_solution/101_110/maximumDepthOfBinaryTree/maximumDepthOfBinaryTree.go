package maximumDepthOfBinaryTree

// 104. Maximum Depth of Binary Tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lHeight := maxDepth(root.Left)
	rHeight := maxDepth(root.Right)
	return max(lHeight, rHeight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
