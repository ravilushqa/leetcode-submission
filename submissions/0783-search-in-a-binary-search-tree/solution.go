/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }
    
    if root.Val == val {
        return root
    }
    
    leftRes := searchBST(root.Left, val)
    if leftRes != nil {
        return leftRes
    }
    
    return searchBST(root.Right, val)
}
