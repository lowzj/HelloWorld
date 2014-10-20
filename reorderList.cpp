#include <iostream>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};
       
class Solution {
public:
    void reverseList(ListNode* pre) {
        if (!pre) return;
        ListNode* cur = pre->next;
        pre->next = NULL;
        while (cur) {
            ListNode* next = cur->next;
            cur->next = pre->next;
            pre->next = cur;
            cur = next;
        }
    }
    void reorderList(ListNode *head) {
        if (!head || !head->next || !head->next->next)
            return;

        int len = 0;
        for (ListNode* p=head; p!=NULL; p=p->next, ++len);

        ListNode* pre = head;
        for (int i=(len-1)/2; i>0; --i, pre=pre->next);

        reverseList(pre);

        // merge
        ListNode* p1 = head;
        ListNode* p2 = pre->next;
        pre->next = NULL;
        while (p1 && p2) {
            ListNode* tmp1 = p1->next;
            ListNode* tmp2 = p2->next;

            p1->next = p2;
            p2->next = tmp1;

            p1 = tmp1;
            p2 = tmp2;
        }
    }
};

int main(void) {
    int a[] = {4,19,14,5,-3,1,8,5,11};
    ListNode* head = NULL;
    ListNode* tmp = head;
    for (int i = 0; i < sizeof(a)/sizeof(int); ++i) {
        if (!head) {
            head = new ListNode(a[i]);
            tmp = head;
        } else {
            tmp->next = new ListNode(a[i]);
            tmp = tmp->next;
        }
    }
    tmp->next = NULL;
    for (ListNode* p = head; p; p = p->next) {
        cout << p->val << " ";
    }
    cout << endl;
    Solution s;
    s.reorderList(head);
    for (ListNode* p = head; p; p = p->next) {
        cout << p->val << " ";
    }
    cout << endl;
    return 0;
}
