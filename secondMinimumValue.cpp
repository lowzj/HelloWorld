#include <iostream>


using namespace std;

// https://leetcode.com/contest/leetcode-weekly-contest-48/problems/second-minimum-node-in-a-binary-tree/
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
class Solution {
public:
    int find(TreeNode* root, int smallest) {
        if (root == NULL) {
            return -1;
        }
        if (root->val > smallest) return root->val;
        int l = find(root->left, smallest);
        int r = find(root->right, smallest);
        if (l > smallest && (r == -1 || r >= l)) return l;
        if (r > smallest && (l == -1 || l >= r)) return r;
        return -1;
    }
    int findSecondMinimumValue(TreeNode* root) {
        return find(root, root->val);
    }
};

int main(void) {
    TreeNode* r = new TreeNode(2);
    r->left = new TreeNode(2);
    r->right = new TreeNode(5);
    r->right->left = new TreeNode(5);
    r->right->right = new TreeNode(7);
    Solution s;
    cout << s.findSecondMinimumValue(r) << endl;
    return 0;
}
