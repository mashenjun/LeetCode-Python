package binaryTreePaths

import (
	"strconv"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 257. Binary Tree Paths

// Given a binary tree, return all root-to-leaf paths.
// Note: A leaf is a node with no children.
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-tree-paths
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：dfs+临时数组，记录走过的路径
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	stack := []*TreeNode{root}
	paths := []string{strconv.Itoa(root.Val)}
	ret := make([]string, 0)
	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		currPath := paths[len(paths)-1]
		if curr.Right == nil && curr.Left == nil {
			ret = append(ret, currPath)
		}
		stack = stack[:len(stack)-1]
		paths = paths[:len(paths)-1]
		if curr.Right != nil {
			stack = append(stack, curr.Right)
			newPath := currPath +"->"+strconv.Itoa(curr.Right.Val)
			paths = append(paths, newPath)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
			newPath := currPath +"->"+strconv.Itoa(curr.Left.Val)
			paths = append(paths, newPath)
		}
	}
	return ret
}


func TestBinaryTreePath(t *testing.T) {
	root := &TreeNode{
		Val:1,
		Left:&TreeNode{
			Val:   2,
			Left: &TreeNode{
				Val:   5,
			},
			Right:&TreeNode{
				Val:   6,
			},
		},
		Right:&TreeNode{
			Val:   3,
		},
	}
	t.Log(binaryTreePaths(root))
}
