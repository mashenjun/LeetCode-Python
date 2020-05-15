package pointOffer

import "testing"

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder, 0, 0, len(inorder)-1)
}
// 思路：递归，每次递归需要的参数为 需要构建的子树的根节点在preorder中的下标，子树在inorder中的左边界和右边界
func build(preorder []int, inorder []int, idx, left, right int) *TreeNode{
	if left > right {
		return nil
	}
	// find the pivot in inorder
	pivot := 0
	for i:= 0; i < len(preorder); i++ {
		if preorder[idx] == inorder[i] {
			pivot = i
			break
		}
	}

	root := &TreeNode{
		Val: preorder[idx],
		Left: build(preorder, inorder, idx+1, left, pivot-1),
		Right: build(preorder, inorder, idx+pivot-left+1, pivot+1, right),
	}
	return root
}

func TestBuildTree(t *testing.T) {
	preOrder := []int{3,9,20,15,7}
	inOrder := []int{9,3,15,20,7}

	root := buildTree(preOrder, inOrder)
	t.Log(root.Val)
}
