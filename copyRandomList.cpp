#include <iostream>
#include <map>

using namespace std;

struct RandomListNode {
    int label;
    RandomListNode *next, *random;
    RandomListNode(int x) : label(x), next(NULL), random(NULL) {}
};
       
class Solution {
public:
    RandomListNode *copyRandomList(RandomListNode *head) {
        map<RandomListNode*, RandomListNode*> m;
        
        for (RandomListNode *tmp=head; tmp; tmp=tmp->next) {
            RandomListNode *copy = new RandomListNode(tmp->label);
            m.insert(make_pair(tmp, copy));
        }
        for (RandomListNode *tmp=head; tmp; tmp=tmp->next) {
            RandomListNode *copy = m[tmp];
            if (tmp->next) 
                copy->next = m[tmp->next];
            if (tmp->random)
                copy->random = m[tmp->random];
        }
        if (m.empty()) return NULL;
        return m[head];
    }
    RandomListNode *copyRandomList2(RandomListNode *head) {
        if (!head) return NULL;
        
        for (RandomListNode *p=head; p;) {
            RandomListNode *tmp = new RandomListNode(p->label);
            tmp->next = p->next;
            p->next = tmp;
            p = tmp->next;
        }

        for (RandomListNode *p=head; p && p->next; p=p->next->next) {
            p->next->random = p->random ? p->random->next : NULL;
        }

        RandomListNode *copyHead = head->next;
        RandomListNode *p1 = head;
        RandomListNode *p2 = copyHead;

        while (p2) {
            p1->next = p1->next->next;
            if (p2->next)
                p2->next = p2->next->next;
            p1 = p1->next;
            p2 = p2->next;
        }

        return copyHead;
    }
};

int main(void) {
    Solution s;
    s.copyRandomList(NULL);
    return 0;
}

