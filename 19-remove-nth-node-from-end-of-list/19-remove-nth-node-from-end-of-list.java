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
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode  dum = new ListNode(-1);
        dum.next = head;
        ListNode show = dum;   //满指针
        ListNode fort = dum;   //快指针
        
        
        while(fort != null && n-- > 0) {
            fort = fort.next;
        }
        ListNode pre = null;
        while(fort != null) {
            pre = show;
            show = show.next;
            fort = fort.next;
        }
        
        pre.next = show.next;
        show.next = null;
        return dum.next;
    
    }
}