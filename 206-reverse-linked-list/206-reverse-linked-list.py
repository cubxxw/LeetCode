# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        pre, cur = None, head
        while cur:
            tmp = cur.next   #保存当前结点
            cur.next = pre
            pre = cur
            cur = tmp    #当前结点是之前保存的结点  后移一位
        return pre   #返回的是前一个结点        
    