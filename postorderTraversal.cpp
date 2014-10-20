#include <iostream>
#include <vector>
#include <stack>
#include <algorithm>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

// postorder:   left,right,root
// preorder:    root,left,right
// inorder:     left,root,right
class Solution {
public:
    vector<int> postorderTraversal(TreeNode *root) {
        vector<int> ret;
        stack<TreeNode*> s;

        if (root) {
            s.push(root);
        }
        while (!s.empty()) {
            TreeNode* tmp = s.top();
            s.pop();
            ret.push_back(tmp->val);
            if (tmp->left) {
                s.push(tmp->left);
            }
            if (tmp->right) {
                s.push(tmp->right);
            }
        }
        reverse(ret.begin(), ret.end());
        return ret;
    }
};

int main(void) {
    TreeNode* root = new TreeNode(1);
    TreeNode* left = new TreeNode(3);
    TreeNode* right = new TreeNode(2);
    root->right = right;
    root->left = left;

    Solution s;
    vector<int> v = s.postorderTraversal(root);
    for (int i = 0; i < v.size(); ++i) {
        cout << v[i] << " ";
    }
    cout << endl;
    return 0;
}
