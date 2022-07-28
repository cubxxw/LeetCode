// *
//  * Definition for singly-linked list.
//  * public class ListNode {
//  *     int val;
//  *     ListNode next;
//  *     ListNode() {}
//  *     ListNode(int val) { this.val = val; }
//  *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
//  * }

class Solution {
    public static ListNode reverseKGroup(ListNode head, int k) {
        ListNode start = head;
        ListNode end = getKGroupEnd(start,k);   //数到k的点返回给end
        if(end == null) {
            //没凑齐，直接返回
            return head;
        }
        head = end;      //凑齐了，就把最后面的给头,头不变
        reverse(start,end);   //反转
        ListNode lastEnd = start;   //上一组的结尾结点
        while(lastEnd.next != null) {
            start = lastEnd.next; //新结点
            end = getKGroupEnd(start,k);  //最后一个结点调整
            if(end == null) {
                return head;
            }
            reverse(start,end);
            lastEnd.next = end;
            lastEnd = start;
        }
        return head;  //正好是整数倍
    }
    
    public static void reverse(ListNode start, ListNode end) {
        //之前的分区最后一个结点，s前一个结点不用关心,也看不到，指向null就行
        end = end.next;
        ListNode pre = null;
        ListNode cur = start;
        ListNode next = null;   //cur当前来到start位置
        while(cur != end) {  //没到最后面，需要走到最后的时候往回指向，和上面单链表反转操作一样
            next = cur.next;  
            cur.next = pre;
            pre = cur;
            cur = next;
        }
        start.next = end;    //原来头结点连接原来尾结点 --> next
    }
    
    public static ListNode getKGroupEnd(ListNode start, int k) {
        while(--k != 0 && start != null) {
            //k = 5  k = 0 end   start = 123456 , return 5    而且不够k返回null
            start = start.next;
        }
        return start;
    }
    
}