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
    public boolean isPalindrome(ListNode head) {
        List<Integer> list = new ArrayList<>();
        ListNode iterator = head;
        ListNode helper = null;
        while (iterator != null) {
            list.add(iterator.val);
            ListNode nextNode = iterator.next;
            // point to back
            iterator.next = helper;
            // move helper
            helper = iterator;
            // move iterator to the next node
            iterator = nextNode;
        }

        iterator = helper;
        int i = 0;
        while (iterator != null) {
            if (list.get(i) != iterator.val) {
                System.out.println(i);
                System.out.println(list.get(i));
                System.out.println(iterator.val);
                return false;
            }
            iterator = iterator.next;
            i++;
        }

        return true;
    }
}
