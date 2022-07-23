/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode  //定义一个链表
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev   //改变链表
        prev = curr  //
        curr = next
    }
    return prev
}
