package pointOffer

import (
	"fmt"
	"testing"
)

var tree *TreeNode

func init() {
	tree = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left:&TreeNode{
				Val: 1,
				Right: &TreeNode{Val: 2},
			},
			Right: &TreeNode{
				Val: 4,
			},

		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{
				Val: 8,
				Right:&TreeNode{Val: 9},
			},
		},
	}
}

func TestBinarySearchTreeConvert(t *testing.T) {
	head ,tail:= inorderIteration(tree)
	fmt.Printf("head %v\n", head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Printf("tail %v\n", tail)
	for tail != nil {
		fmt.Println(tail.Val)
		tail = tail.Prev
	}
}

