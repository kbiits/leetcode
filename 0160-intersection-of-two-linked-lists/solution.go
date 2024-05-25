/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// var memory = map[*ListNode]bool{}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    f, s := headA, headB

    for f != s {
        if f == nil {
            f = headB
        }
        if s == nil {
            s = headA
        }
        if f == s {
            return f
        }
        f, s = f.Next, s.Next
    }

    return f
}

// 2 + 3
// 5 + 3
