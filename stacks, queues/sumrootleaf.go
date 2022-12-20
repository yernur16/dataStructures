package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Stack is a slice of *TreeNode pointers
type Stack []*TreeNode

// Push adds a new element to the top of the stack
func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() *TreeNode {
	if len(*s) == 0 {
		return nil
	}

	node := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return node
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Initialize stack and result
	stack := Stack{root}
	res := 0

	for len(stack) > 0 {
		// Pop top element from stack
		node := stack.Pop()

		// If node is a leaf, add its value to the result
		if node.Left == nil && node.Right == nil {
			res += node.Val
		} else {
			// Push left and right children onto the stack
			if node.Left != nil {
				node.Left.Val = node.Val*10 + node.Left.Val
				stack.Push(node.Left)
			}
			if node.Right != nil {
				node.Right.Val = node.Val*10 + node.Right.Val
				stack.Push(node.Right)
			}
		}
	}

	return res
}
