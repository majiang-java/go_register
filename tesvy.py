class Solution:
    def mergeTwoLists(self, list1, list2):
        """

        """
        merged = ListNode(None)
        result = merged

        while list1 and list2:
            if list1.val<=list2.val:
                smaller = list1.val
                list1 = list1.next
            else:
                smaller = list2.val
                list2 = list2.next
        merged.next = ListNode(smaller)
        merged = merged.next

        while list1:
            merged.next = ListNode(list1.val)
            list1 = list1.next
            merged = merged.next
        while list2:
            merged.next = ListNode(list2.val)
            list2 = list2.next
            merged = merged.next
        return result.next


