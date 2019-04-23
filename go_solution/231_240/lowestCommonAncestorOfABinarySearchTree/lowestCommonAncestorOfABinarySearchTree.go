package lowestCommonAncestorOfABinarySearchTree

// 235. Lowest Common Ancestor of a Binary Search Tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：利用搜索二叉树的特性
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val < p.Val && root.Val < q.Val { // 同时小于 往右走
			root = root.Right
		} else if root.Val > p.Val && root.Val > q.Val { // 同时大于 往左走
			root = root.Left
		} else {
			return root
		}
	}
	return root
}
