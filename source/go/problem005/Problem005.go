package problem005

import "fmt"

type TreeNode struct {
	Val int
	Lchild * TreeNode
	Rchild * TreeNode
}

func FindValue(s []int, val int) int {
	for key, value := range s{
		if value == val{
			return key
		}
	}
	return -1
}

func PrintTree(node *TreeNode, level int) {
	if node == nil {
		return
    }

	format := ""
	for i := 0; i < level; i++ {
		format += "\t"
	}
	format += "----[ "
	level++
	PrintTree(node.Lchild, level)
	fmt.Printf(format+"%d\n", node.Val)
	PrintTree(node.Rchild, level)
}

func ConstructTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}
	index := FindValue(inorder, preorder[0])
	left := inorder[0:index]
	right := inorder[index+1:]
	root := TreeNode{Val:preorder[0]}
	root.Lchild = ConstructTree(preorder[1:1+len(left)], left)
	root.Rchild = ConstructTree(preorder[1+len(left):], right)
	return &root
}
