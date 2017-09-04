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
    TreeNode* trimBST(TreeNode* root, int L, int R) {
        if (root == NULL) return NULL;
        TreeNode* trim = NULL;
        if (root->val >= L && root->val <= R) {
            trim = root;
            if (root->left != NULL) trim->left = trimBST(root->left, L, R);
            if (root->right != NULL) trim->right = trimBST(root->right, L, R);
        } else if (root->val < L) {
            return trimBST(root->right, L, R);
        } else if (root->val > R) {
            return trimBST(root->left, L, R);
        }
        return trim;
    }
};

void print(TreeNode* t) {
    if (t == NULL) return;
    if (t->left != NULL) print(t->left);
    cout << t->val << " ";
    if (t->right != NULL) print(t->right);
}

int main(void) {
    TreeNode* r = new TreeNode(3);
    r->left = new TreeNode(0);
    r->right = new TreeNode(4);
    r->left->right = new TreeNode(2);
    r->left->right->left = new TreeNode(1);
    Solution s;
    TreeNode* t = s.trimBST(r, 1, 3);
    print(t);
    cout << endl;
    return 0;
}
