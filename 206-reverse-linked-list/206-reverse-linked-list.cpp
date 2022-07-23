/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* reverseList(ListNode* head) {
    ListNode* pre = nullptr;
    ListNode* cur = head;
    while(cur != nullptr){
        ListNode* tmp = cur->next;    //保存下一个结点
        cur->next = pre;    //指向前一个结点
        pre = cur;    //保存后续的结点
        cur = tmp;   //当前结点是之前保存的哪个结点
        }
    return pre;
    }        
};