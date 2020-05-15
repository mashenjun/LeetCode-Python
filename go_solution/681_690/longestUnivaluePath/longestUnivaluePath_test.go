package longestUnivaluePath

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：递归
// 终止条件：
// 每次递归的返回值：
// 单次递归的工作：
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ret = 0
	_ = find(root.Left, root.Val, &ret)
	_ = find(root.Right, root.Val, &ret)
	return ret
}

func find(root *TreeNode, val int, ret *int) int {
	if root == nil {
		return 0
	}
	l := find(root.Left, root.Val, ret)
	r := find(root.Right, root.Val, ret)
	m := max(l, r)
	*ret = max(*ret, r+l)

	if root.Val != val {
		return 0
	}
	return m + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestLongestUnivaluePath(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   5,
			},
		},
	}

	t.Log(longestUnivaluePath(root))
}
