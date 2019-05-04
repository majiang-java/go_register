class ListNode:
    def __init__(self, x) :
        self.val = x
        self.next = None

class Solution:
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        l = []
        for head_node in lists:
            cur = head_node
            while cur:
                l.append(cur.val)
                cur = cur.next

        l.sort()
        dummy_head = ListNode(0)
        for n in l:
            cur.next = ListNode(n)
            cur = cur.next
        return dummy_head.next        