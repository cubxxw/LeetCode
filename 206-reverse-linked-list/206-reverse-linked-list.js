/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var reverseList = function(head) {
    var prev = null  //定义一个链表
    var curr = head
    while(curr) {
        const next = curr.next;
        curr.next = prev;   //改变链表
        prev = curr;  
        curr = next;
    }
    return prev;
};