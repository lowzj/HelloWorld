#include <iostream>
using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
        
const int INFINT = 0x80000000;
class Solution {
    int maxSum_;
public:
    int calMaxPathSum(TreeNode *root) {
        if (!root) return INFINT;
        int left = INFINT;
        int right = INFINT;
        if (root->left) {
            left = calMaxPathSum(root->left);
        }
        if (root->right) {
            right = calMaxPathSum(root->right);
        }
        int maxsum = std::max(left, right);
        maxsum = maxsum==INFINT ? root->val : max(root->val, root->val+maxsum);
        maxSum_ = max(maxSum_, maxsum);
        if (left != INFINT && right != INFINT)
            maxSum_ = max(maxSum_, root->val+left+right);
        cout << maxSum_ << endl;

        return maxsum;
    }
    int maxPathSum(TreeNode *root) {
        maxSum_ = INFINT;
        calMaxPathSum(root);
        return maxSum_;
    }
};

int main(void) {
    TreeNode *root = new TreeNode(1);
    root->left = new TreeNode(2);
    root->right = new TreeNode(3);
    Solution s;
    cout << s.maxPathSum(root) << endl;
    return 0;
}

