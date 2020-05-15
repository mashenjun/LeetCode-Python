package minimumDistanceBetweenBSTNodes

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	var ret int = 1<<63 -1
	var prev = 0
	findDFS(root, &ret, &prev)
	return ret
}

func findDFS(root *TreeNode, retPtr *int, prev *int){
	// 终止条件
	if root == nil {
		return
	}
	findDFS(root.Left, retPtr, prev)
	if *prev != 0 {
		*retPtr = min(*retPtr, root.Val - *prev)
	}
	*prev = root.Val
	findDFS(root.Right, retPtr, prev)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMinDiffInBST(t *testing.T) {
	root := &TreeNode{
		Val:   27,
		Right:  &TreeNode{
			Val:   34,
			Right:&TreeNode{
				Val: 58,
				Left: &TreeNode{
					Val: 50,
					Left:&TreeNode{
						Val:   44,
					},
				},
			},
		},
	}

	t.Log(minDiffInBST(root))
}
