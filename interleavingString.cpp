#include <iostream>
using namespace std;

class Solution {
    private:
        int len1;
        int len2;
        int len3;
        int *m;
    public:
        int solve(int i, int j, string& s1, string& s2, string& s3) {
            cout << i << " " << j << endl;
            if (i+j == len3) return 1;
            cout << i << " " << j << "a" << m[i*len1+j] << endl;
            if (m[i*len1+j] != 0) return m[i*len1+j];
            cout << i << " " << j << "b" << endl;

            int tmp = -1;
            if (i < len1 && s1[i] == s3[i+j]) {
                tmp = solve(i+1, j, s1, s2, s3);
            }
            if (tmp != 1 && j < len2 && s2[j] == s3[i+j]) {
                tmp = solve(i, j+1, s1, s3, s3);
            }
            m[i*len1+j] = tmp;
            cout << i << " " << j << "c" << tmp << endl;
            return m[i*len1+j];
        }
        bool isInterleave(string s1, string s2, string s3) {
            len1 = s1.size();
            len2 = s2.size();
            len3 = s3.size();
            m = new int((len1+1)*(len2+1));
            memset(m, 0, (len1+1)*(len2+1));

            return (len1+len2 == len3) && solve(0, 0, s1, s2, s3) == 1;
        }
};
int main(void) {
    string s1("aabd");
    string s2("adbc");
    string s3("aabdabdc");
    Solution s;
    cout << s.isInterleave(s1, s2, s3) << endl;
    return 0;
}
