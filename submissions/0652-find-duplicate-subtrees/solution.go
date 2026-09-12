/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    res := []*TreeNode{}
    dublicatesCount := make(map[string]int) 
    var dfs func(node *TreeNode) string

    dfs = func(node *TreeNode) string {
        if node == nil {
            return "#"
        }
        
        key := fmt.Sprintf("%d_%s_%s", node.Val, dfs(node.Left), dfs(node.Right))
        dublicatesCount[key]++

        if dublicatesCount[key] == 2 {
            res = append(res, node)
        }

        return key
    }

    dfs(root)

    return res
}



