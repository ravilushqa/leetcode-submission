/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var dfs func(node *TreeNode, res []int) []int

    dfs = func(node *TreeNode, res []int) []int{
        if node == nil {
            return res
        }

        res = dfs(node.Left, res)
        res = append(res, node.Val)
        res = dfs(node.Right, res)

        return res
    }

    return dfs(root, nil)
}
