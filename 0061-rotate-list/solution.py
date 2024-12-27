# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        if k == 0 or head is None:
            return head
        
        def countNodes():
            cur = head
            count = 0
            while cur is not None:
                count += 1
                cur = cur.next
            return count

        count = countNodes()
        last_idx = k % count
        if last_idx == 0:
            return head

        skip = count - last_idx
        cur = head
        last_node = None
        for i in range(skip):
            if i == skip - 1:
                last_node = cur
            cur = cur.next
        
        # fix the pointer
        cur2 = cur
        while cur2.next is not None:
            cur2 = cur2.next
        cur2.next = head
        last_node.next = None
        return cur
