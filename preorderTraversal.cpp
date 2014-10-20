#include <iostream>
#include <vector>
#include <stack>

using namespace std;


struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    vector<int> preorderTraversal(TreeNode *root) {
        vector<int> ret;
        stack<TreeNode*> s;
        if (root)
            s.push(root);
        while (!s.empty()) {
            TreeNode* tmp = s.top();
            ret.push_back(tmp->val);
            s.pop();

            if (tmp->right)
                s.push(tmp->right);
            if (tmp->left)
                s.push(tmp->left);
        }
        return ret;
    }
};

int main(void) {
    return 0;
}
