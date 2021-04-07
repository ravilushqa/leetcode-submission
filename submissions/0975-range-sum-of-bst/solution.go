/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	res := 0

	if root.Val >= low && root.Val <= high {
		res+=root.Val
	}
	if root.Val >= low && root.Val <= high {
		return res + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	}

	if root.Val <= low {
		return res + rangeSumBST(root.Right, low, high)
	}

	if root.Val >= high {
		return res + rangeSumBST(root.Left, low, high)
	}

	return res
}
