/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    return maxDepthRecursive(root, 0)
}

func maxDepthRecursive(node *TreeNode, res int) int {
    if node == nil {
        return res
    }

    return 1 + max(maxDepthRecursive(node.Left, res), maxDepthRecursive(node.Right, res))
}
