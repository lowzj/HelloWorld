#include <iostream>
#include <vector>
#include <queue>
#include <unordered_map>
#include <unordered_set>

using namespace std;

struct UndirectedGraphNode {
    int label;
    vector<UndirectedGraphNode *> neighbors;
    UndirectedGraphNode(int x) : label(x) {};
};
       
class Solution {
public:
    typedef unordered_map<UndirectedGraphNode*, UndirectedGraphNode*> GraphMap;
    typedef GraphMap::iterator GraphMapItr;
    UndirectedGraphNode *cloneGraph(UndirectedGraphNode *node) {
        GraphMap m;
        queue<UndirectedGraphNode*> q;
        
        if (!node) return NULL;

        q.push(node);
        while (!q.empty()) {
            UndirectedGraphNode *tmp = q.front();
            q.pop();
            if (m.find(tmp) != m.end()) {
                continue;
            }
            UndirectedGraphNode *copy = new UndirectedGraphNode(tmp->label);
            m.insert(make_pair(tmp, copy));

            for (int i = tmp->neighbors.size()-1; i >= 0; --i) {
                if (m.find(tmp->neighbors[i]) == m.end()) {
                    q.push(tmp->neighbors[i]);
                }
            }
        }

        q.push(node);
        unordered_set<UndirectedGraphNode*> mark;
        while (!q.empty()) {
            UndirectedGraphNode *tmp = q.front();
            q.pop();
            if (mark.find(tmp) != mark.end()) {
                continue;
            }
            mark.insert(tmp);

            UndirectedGraphNode *copy = m[tmp];
            for (int i = tmp->neighbors.size()-1; i >= 0; --i) {
                copy->neighbors.push_back(m[tmp->neighbors[i]]);
                if (mark.find(tmp->neighbors[i]) == mark.end()) {
                    q.push(tmp->neighbors[i]);
                }
            }
        }

        return m[node];
    }
};

int main(void) {
    UndirectedGraphNode *node0 = new UndirectedGraphNode(0);
    UndirectedGraphNode *node1 = new UndirectedGraphNode(1);
    UndirectedGraphNode *node2 = new UndirectedGraphNode(2);
    UndirectedGraphNode *node3 = new UndirectedGraphNode(3);
    UndirectedGraphNode *node4 = new UndirectedGraphNode(4);
    UndirectedGraphNode *node5 = new UndirectedGraphNode(5);

    node0->neighbors.push_back(node1);
    node0->neighbors.push_back(node5);
    node1->neighbors.push_back(node2);
    node1->neighbors.push_back(node5);
    node2->neighbors.push_back(node3);
    node3->neighbors.push_back(node4);
    node3->neighbors.push_back(node4);
    node4->neighbors.push_back(node5);
    node4->neighbors.push_back(node5);

    Solution s;
    s.cloneGraph(node0);

    return 0;
}

