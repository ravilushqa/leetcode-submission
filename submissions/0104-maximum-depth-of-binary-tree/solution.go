/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    return dfs(root, 0)
}

func dfs(node *TreeNode, height int) int {
    if node == nil {
        return height
    }

    height++

    return max(dfs(node.Left, height), dfs(node.Right, height))
}
