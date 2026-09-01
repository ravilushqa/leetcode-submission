/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return dfs(root) >= 0 
}

func dfs(node *TreeNode) int {
    if node == nil {
        return 0
    }

    lh := dfs(node.Left)
    rh := dfs(node.Right)

    if lh == -1 || rh == -1 {
        return -1
    }

    if abs(lh-rh) >= 2 {
        return -1
    }


    return 1 + max(lh, rh)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}
