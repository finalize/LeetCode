package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	var sum int

	if root != nil {
		addIntegers(root, &sum, L, R)
	}
	fmt.Println(sum)
	return sum
}

func addIntegers(node *TreeNode, sum *int, L int, R int) {

	if node.Val >= L && node.Val <= R {
		*sum += node.Val
	}

	if node.Left != nil {
		addIntegers(node.Left, sum, L, R)
	}
	if node.Right != nil {
		addIntegers(node.Right, sum, L, R)
	}
}

func main() {

}
