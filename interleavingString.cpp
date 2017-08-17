#include <iostream>
using namespace std;

class Solution {
    private:
        int len1;
        int len2;
        int len3;
        int **m;
    public:
        int solve(int r, int c, string& s1, string& s2, string& s3) {
            if (r+c == len3) return 1;
            if (m[r][c] != 0) return m[r][c];

            int tmp = -1;
            if (r < len1 && s1[r] == s3[r+c]) {
                tmp = solve(r+1, c, s1, s2, s3);
            }
            if (tmp != 1 && c < len2 && s2[c] == s3[r+c]) {
                tmp = solve(r, c+1, s1, s2, s3);
            }
            m[r][c] = tmp;
            return m[r][c];
        }
        bool isInterleave(string s1, string s2, string s3) {
            len1 = s1.size();
            len2 = s2.size();
            len3 = s3.size();
            m = new int*[len1+1];
            for (int i = 0; i <= len1; ++i) {
                m[i] = new int[len2+1];
                for (int j = 0; j <= len2; ++j) m[i][j] = 0;
            }

            return (len1+len2 == len3) && solve(0, 0, s1, s2, s3) == 1;
        }
};
int main(void) {
    string s1("aabcc");
    string s2("dbbca");
    string s3("aadbbcbcac");
    Solution s;
    cout << s.isInterleave(s1, s2, s3) << endl;
    return 0;
}
