package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := []*TreeNode{root}
	result := 0

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Left == nil && node.Right == nil {
			result += node.Val
		} else {
			if node.Left != nil {
				node.Left.Val = node.Val*2 + node.Left.Val
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = node.Val*2 + node.Right.Val
				stack = append(stack, node.Right)
			}
		}

	}
	return result
}
