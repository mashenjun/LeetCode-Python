package rangeSumOfBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 终止条件：root为nil
// 返回值：以当前root为根的子树，在[L,R]之间的节点的值的和
// 单次递归的逻辑：
// 判断 root.Val 小于L -> return rangeSumBST(root.Right, L,R);
//     root.Val  大于R -> reutrn rangeSumBST(root.Left, L,R);
//     root.Val 在区间内 -> return root.Val + rangeSumBST(root.Right, L,R) + rangeSumBST(root.Left, L,R)
//
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	l := rangeSumBST(root.Left, L, R)
	r := rangeSumBST(root.Right, L, R)
	if root.Val < L {
		return r
	}
	if root.Val > R {
		return l
	}
	return root.Val + l + r
}
