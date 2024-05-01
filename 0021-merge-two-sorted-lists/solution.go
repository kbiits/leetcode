/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    newHead := &ListNode{}
    ptr := newHead


    for list1 != nil || list2 != nil {
        if (list1 != nil && list2 != nil && list1.Val <= list2.Val) || list2 == nil {
            ptr.Next = list1
            list1 = list1.Next
        } else if list2 != nil {
            ptr.Next = list2
            list2 = list2.Next
        }

        ptr = ptr.Next
    }

    return newHead.Next
}

