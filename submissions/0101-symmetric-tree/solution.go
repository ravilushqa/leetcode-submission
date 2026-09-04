/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return IsSymetricR(root.Left, root.Right)
}

func IsSymetricR(n1, n2 *TreeNode) bool {
    if n1 == nil && n2 == nil {
        return true
    }

    if n1 == nil || n2 == nil {
        return false
    } 

    return n1.Val == n2.Val && IsSymetricR(n1.Left, n2.Right) && IsSymetricR(n1.Right, n2.Left)
}

