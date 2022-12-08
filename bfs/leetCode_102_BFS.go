// Definition for a binary tree node.
//
//	type TreeNode struct {
//		Val   int
//		Left  *TreeNode
//		Right *TreeNode
//	}
package main

type NodeWithLevel struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	var tmp = []int{}
	var res = [][]int{}
	//создаем два массива. 1 обычный и другой двумерный
	//потом аппендим обычный массив в двумерный массив
	if root == nil {
		return res
	}
	q := []NodeWithLevel{
		{
			node:  root,
			level: 0,
		},
	}
	res = [][]int{{
		q[0].node.Val,
	}}

	var ct = 0

	for len(q) > 0 {
		vertex := q[0]
		node, level := vertex.node, vertex.level
		q = q[1:]

		if ct < level {
			ct = level
			res = append(res, tmp)
			tmp = []int{}
		}

		if node.Left != nil {
			leftNode := NodeWithLevel{
				node:  node.Left,
				level: level + 1,
			}
			tmp = append(tmp, leftNode.node.Val)

			q = append(q, leftNode)
		}
		if node.Right != nil {
			rightNode := NodeWithLevel{
				node:  node.Right,
				level: level + 1,
			}
			tmp = append(tmp, rightNode.node.Val)

			q = append(q, rightNode)
		}

	}
	return res
}

// func main() {
// 	root := []TreeNode{3, 9, 20, null, null, 15, 7}
// 	fmt.Println(levelOrder(root))
// }

// func main() {
// 	root := []TreeNode{3, 9, 20, null, null, 15, 7}
// 	fmt.Println(levelOrder(root))
// }
