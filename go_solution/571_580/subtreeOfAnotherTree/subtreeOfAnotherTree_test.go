package subtreeOfAnotherTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 572. Subtree of Another Tree
// Given two non-empty binary trees s and t, check whether tree t has exactly the same structure and node values with a subtree of s.
// A subtree of s is a tree consists of a node in s and all of this node's descendants. The tree s could also be considered as a subtree of itself.
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/subtree-of-another-tree
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：结合IsSameTree的解法，递归判断是否是相似的树
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	return isSame(s, t) || isSubtree(s.Right, t) || isSubtree(s.Left, t)
}

func isSame(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val && isSame(l.Left, r.Left) && isSame(l.Right, r.Right)
}
