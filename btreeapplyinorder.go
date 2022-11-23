package main

import "fmt"

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}

func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	BTreeApplyInorder(root, fmt.Println)
}

func BTreeInsertData(root *TreeNode, Data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: Data}
	}
	if Data < root.Data {
		root.Left = BTreeInsertData(root.Left, Data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, Data)
		root.Right.Parent = root
	}
	return root
}
