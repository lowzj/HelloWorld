#include <iostream>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode *insertionSortList(ListNode *head) {
        ListNode* sorted = new ListNode(0);

        ListNode* p = head;
        while (p) {
            ListNode* pos = sorted;
            while (pos->next && pos->next->val < p->val) {
                pos = pos->next;
            }

            ListNode* tmp = p->next;
            p->next = pos->next;
            pos->next = p;
            p = tmp;
        }

        ListNode* sortedHead = sorted->next;
        delete sorted;
        return sortedHead;
    }
};

int main(void) {
    int a[] = {4,19,14,5,-3,1,8,5,11,15};
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
    Solution s;
    head = s.insertionSortList(head);
    for (ListNode* p = head; p; p = p->next) {
        cout << p->val << endl;
    }
    return 0;
}

