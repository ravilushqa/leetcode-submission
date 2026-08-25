/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    return inOrderRecursive(root, nil)
}

func inOrderRecursive(node *TreeNode, res []int) []int {
    if node == nil {
        return res
    }

    res = inOrderRecursive(node.Left, res)
    res = append(res, node.Val)
    res = inOrderRecursive(node.Right, res)

    return res
}


