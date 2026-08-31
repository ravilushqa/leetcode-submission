/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    return dfs(root, nil)
}

func dfs(node *TreeNode, res []int) []int {
    if node == nil {
        return res
    }

    res = dfs(node.Left, res)
    res = append(res, node.Val)
    res = dfs(node.Right, res)

    return res
}
