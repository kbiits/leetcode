/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    } else if list2 == nil {
        return list1
    }

    if list1.Val < list2.Val {
        list1.Next = mergeTwoLists(list2, list1.Next)
        return list1
    }
    list2.Next = mergeTwoLists(list1, list2.Next)
    return list2
}
