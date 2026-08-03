/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    return dfsInOrder(root, nil)
}

func dfsInOrder(node *TreeNode, res []int) []int {
    if node == nil {
        return []int{}
    }

    if node.Left != nil {
        res = dfsInOrder(node.Left, res)
    }

    res = append(res, node.Val)

    if node.Right != nil {
        res = dfsInOrder(node.Right, res)
    }

    return res
}
