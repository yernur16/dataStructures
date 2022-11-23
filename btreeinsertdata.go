package main

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
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
