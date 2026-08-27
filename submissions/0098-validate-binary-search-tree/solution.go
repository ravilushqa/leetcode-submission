/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return valid(root, nil, nil) 
}

func valid(node *TreeNode, lo, hi *int) bool {
    if node == nil {
        return true
    }
    if (lo != nil && *lo >= node.Val) || (hi != nil && *hi <= node.Val) {
        return false
    }

    return valid(node.Right, &node.Val, hi) && valid(node.Left, lo, &node.Val)
}

