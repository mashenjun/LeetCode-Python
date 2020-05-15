package minimumDepthOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	ldepth := minDepth(root.Left)
	rdepth := minDepth(root.Right)
	if root.Left == nil || root.Right == nil {
		return ldepth + rdepth + 1
	}
	return min(ldepth, rdepth) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
