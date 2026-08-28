/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    node := root
    for {
        if node.Val > p.Val && node.Val > q.Val {
            node = node.Left
            continue
        }
        
        if node.Val < p.Val && node.Val < q.Val {
            node = node.Right
            continue

        }
        
        return node
    }
}
