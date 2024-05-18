/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    root := head
    for root != nil {
        firstNode := root
        secondNode := root.Next
        if secondNode != nil {
            gcdVal := gcd(firstNode.Val, secondNode.Val)
            newNode := ListNode{gcdVal, secondNode}
            firstNode.Next = &newNode
            root = secondNode
        } else {
            root = root.Next
        }
    }

    return head
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    }

    return gcd(b, a % b)
}
