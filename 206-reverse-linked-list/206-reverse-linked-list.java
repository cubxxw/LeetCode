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
    ListNode pre = null;
    ListNode cur = head;
    while(cur != null){
        ListNode tmp = cur.next;    //保存下一个结点
        cur.next = pre;    //指向前一个结点
        pre = cur;    //保存后续的结点
        cur = tmp;   //当前结点是之前保存的哪个结点
        }
    return pre;
    }
}