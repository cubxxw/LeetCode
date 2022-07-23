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
     if(head == null || head.next == null){
        return head;
    }
    ListNode listnode = reverseList(head.next);   
    //指向下一个结点的值， 最后返回的是5
    
    head.next.next = head;
    head.next = null;
    return listnode;
    }
}