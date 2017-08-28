#include <iostream>
#include <stack>
#include <vector>

using namespace std;

// https://leetcode.com/problems/longest-valid-parentheses/description
class Solution {
public:
    int longestValidParentheses(string s) {
        vector<int> v(s.size(), 0);
        stack<int> st;
        int max = 0;
        int cur = 0;

        for (int i = 0; i < s.length(); ++i) {
            if (s[i] == ')') {
                if (!st.empty()) {
                    v[st.top()] = 1;
                    v[i] = 1;
                    st.pop();
                }
            } else {
                st.push(i);
            }
        }
        for (int i = 0; i < v.size(); ++i) {
            cout << v[i];
            if (v[i]) {
                cur += v[i];
                max = max > cur ? max : cur;
            } else {
                cur = 0;
            }
        }
        cout << endl;
        return max;
    }
};

int main(void) {
    Solution s;
    string str = "))))(()))(()";
    cout << s.longestValidParentheses(str) << endl;
    return 0;
}
