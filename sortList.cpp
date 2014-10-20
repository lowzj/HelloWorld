#include <iostream>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    void partition(ListNode* begin, ListNode* end) {
        if (!begin || begin->next == end || begin == end) return;
        ListNode* p = begin;
        int pval = p->val;
        ListNode* tmp = begin->next;
        while (tmp != end && tmp) {
            if (pval > tmp->val) {
                p->val = tmp->val;
                p = p->next;
                tmp->val = p->val;
            }
            tmp = tmp->next;
        }
        p->val = pval;
        partition(begin, p);
        partition(p->next, end);
    }

    inline void swap(int* a, int* b) {
        if (a == b || !a || !b) return;
        int tmp = *a;
        *a = *b;
        *b = tmp;
    }
    inline ListNode* min(ListNode*& a, ListNode*& b, bool move = true) {
        if (a && b) {
            ListNode* tmp = NULL;
            if (a->val < b->val) {
                tmp = a;
                if (move) a = a->next;
            } else {
                tmp = b;
                if (move) b = b->next;
            }
            return tmp;
        }
        return NULL;
    }
    ListNode* merge(ListNode* begin , ListNode* end, int n) {
        if (n < 2 || !begin) return begin;
        if (n == 2) {
            if (begin->val > begin->next->val) {
                swap(&begin->val, &begin->next->val);
            }
            return begin;
        }

        ListNode* mid = begin;
        int cnt = n / 2;
        while (mid && cnt--) mid = mid->next;

        ListNode* p1 = merge(begin, mid, n/2);
        ListNode* p2 = merge(mid, end, n - n/2);

        ListNode* head = min(p1, p2, true);
        ListNode* p3 = head;
        while (p1 != mid && p2 != end) {
            p3->next = min(p1, p2, true);
            p3 = p3->next;
        }

        ListNode* tmpp = (p1 != mid ? p1 : p2 != end ? p2 : NULL);
        ListNode* tmpend = (p1 != mid ? mid : p2 != end ? end : NULL);
        while (tmpp && tmpp != tmpend) {
            p3->next = tmpp;
            p3 = p3->next;
            tmpp = tmpp->next;
        }
        p3->next = end;
        return head;
    }
    ListNode *sortList(ListNode *head) {
        ListNode* tmp = head;
        int n = 0;
        while (tmp) {
            tmp = tmp->next;
            ++n;
        }
        head = merge(head, NULL, n);
        // partition(head, NULL);
        return head;
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
    head = s.sortList(head);
    for (ListNode* p = head; p; p = p->next) {
        cout << p->val << endl;
    }
    return 0;
}

