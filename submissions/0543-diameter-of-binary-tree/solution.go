/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    _, res := h(root)

    return res
}


func h(node *TreeNode) (height, diameter int) {
    if node == nil {
        return 0, 0
    }

    lh, ld := h(node.Left) 
    rh, rd := h(node.Right)
    

    return 1 + max(lh, rh), max(lh + rh, ld, rd)
}
