/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    ptr := head
    var prev *ListNode = nil
    for ptr != nil {
        next := ptr.Next

        ptr.Next = prev
        
        prev = ptr
        ptr = next
    }

    return prev
}
