#include <iostream>

using namespace std;


struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};
       
#define abs(x) ((x)>0?(x):-(x))
class Solution {
public:
    ListNode *hasCycle(ListNode *head) {
        if (!head) {
            return NULL;
        }
        ListNode *slow = head;
        ListNode *fast = head->next;

        while (fast && slow && fast != slow) {
            slow = slow->next;
            fast = fast->next;
            if (fast) fast = fast->next;
        }
        return fast && slow && fast==slow ? slow : NULL;
    }
    // 相当于求两个相交单向链表的交点
    ListNode *detectCycle(ListNode *head) {
        ListNode *inCycle   = hasCycle(head);
        if (!inCycle) return NULL;

        ListNode *pIn   = NULL;
        ListNode *pOut  = NULL;
        int cycleLen    = 1;
        int slowLen     = 0;
        for (pIn=inCycle->next; pIn!=inCycle; pIn=pIn->next, ++cycleLen);
        for (pOut=head; pOut!=inCycle; pOut=pOut->next, ++slowLen);

        pOut = head;
        pIn = inCycle;
        ListNode **tmp = slowLen>cycleLen ? &pOut : &pIn;
        for (int diff = abs(slowLen-cycleLen); diff > 0; --diff) {
            *tmp = (*tmp)->next;
        }

        while (pIn != pOut) {
            pIn = pIn->next;
            pOut = pOut->next;
        }
        return pIn;
    }

    ListNode *detectCycle2(ListNode *head) {
        ListNode *slow = head;
        ListNode *fast = head;
        
        while (fast && fast->next && slow) {
            slow = slow->next;
            fast = fast->next->next;
            if (slow == fast) {
                fast = head; // head, slow 距离循环起点处有相同的距离
                while (fast != slow) {
                    slow = slow->next;
                    fast = fast->next;
                }
                return slow;
            }
        }
        return NULL;
    }
};

int main(void) {
    return 0;
}
