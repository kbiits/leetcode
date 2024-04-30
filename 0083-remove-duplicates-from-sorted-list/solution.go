/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    for node := head; node != nil; node = node.Next {
        for node.Next != nil && node.Val == node.Next.Val {
            node.Next = node.Next.Next
        }
    }

    return head
}
