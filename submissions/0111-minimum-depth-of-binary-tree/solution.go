/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    var res int
    
    if root == nil {
        return 0 
    }

    queue := []*TreeNode{root}

    for len(queue) > 0 {
        n := len(queue)
        res++
        for i := 1; i<=n; i++ {
            node := queue[0]
            queue = queue[1:]

            if node.Left == nil && node.Right == nil {
                return res
            }
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }

    return 0
}
