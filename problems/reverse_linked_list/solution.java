/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode reverseList(ListNode head) {
        return iterative(head);
    }
    
    public ListNode iterative(ListNode head) {
        ListNode temp = null;
        while(head != null) {
            ListNode next = head.next;
            head.next = temp;
            temp = head;
            head = next;
        }
        return temp;
    }
    
    public ListNode reverseList(ListNode curr, ListNode next) {
        if(next == null) {
            return curr;
        }
        
        ListNode temp = next.next;
        next.next = curr;
        curr = next;
        return reverseList(curr, temp);
    }
    
    
}