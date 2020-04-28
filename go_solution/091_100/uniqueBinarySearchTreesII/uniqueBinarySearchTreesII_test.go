package uniqueBinarySearchTreesII

import "testing"

// 95. Unique Binary Search Trees II
// Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.

// 递归的解法，依次尝试不同的mid值，并把数列分成left部分和right部分,
// left部分和right部分再递归得构造[]*TreeNode
// 把left的slice 和 right 的slice做排列组合，整合成所有可能性
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return toTrees(1, n)
}

func toTrees(left, right int) []*TreeNode{
	if left > right {
		return []*TreeNode{nil}
	}
	if left == right {
		return  []*TreeNode{&TreeNode{
			Val:   left,
			Left:  nil,
			Right: nil,
		}}
	}
	trees := make([]*TreeNode,0)
	for mid := left; mid<=right; mid ++ {
		leftSubTrees:= toTrees(left, mid-1)
		rightSubTrees := toTrees(mid+1, right)
		for _, l := range leftSubTrees{
			for _, r := range rightSubTrees {
				trees = append(trees, &TreeNode{
					Val:   mid,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return trees
}

func TestGenerateTrees(t *testing.T) {
	t.Log(generateTrees(3))
}
