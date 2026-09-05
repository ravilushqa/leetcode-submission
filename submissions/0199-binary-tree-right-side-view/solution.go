/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        n := len(queue)
        for i := 1; i <= n; i++ {
            node := queue[0]
            queue = queue[1:]
            if i == n {
                res = append(res, node.Val)
            }

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }

    return res
}
