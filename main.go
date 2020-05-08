package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var vals []int
	vals = append(vals, root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
	return vals
}

func main() {

}
