/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {    
    el := head
    
    var prev *ListNode
    for el != nil {
        next := el.Next
        el.Next = prev
        prev = el
        el = next
    }

    return prev
}
