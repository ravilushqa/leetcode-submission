/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head

    for fast != nil {
        i := 0
        for fast.Next != nil && i < 2 {
            fast = fast.Next
            i++
        }

        if i == 0 {
            break 
        }

        if i == 1 {
            slow = slow.Next
            break
        }

        if i == 2 {
            slow = slow.Next
        }        
    } 

    return slow
}
