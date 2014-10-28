#include <iostream>
#include <string>
#include <vector>
#include <unordered_set>

using namespace std;

class Solution {
public:
    bool wordBreak(string s, unordered_set<string> &dict) {
        int len = s.size();
        int mark[len];
        int pos = 0;

        for (int i = 0; i < len; ++i) {
            mark[i] = -1;
            for (int j = pos; j >= 0; --j) {
                if (dict.end() != dict.find(s.substr(mark[j]+1, i-mark[j]))) {
                    mark[pos++] = i;
                    break;
                }
            }
        }
        if (pos) return mark[pos-1] == len - 1;
        return false;
    }
};

int main(void) {
    string str = "leetcode";
    unordered_set<string> dict;
    dict.insert("leet");
    dict.insert("code");

    Solution s;
    cout << s.wordBreak(str, dict) << endl;
    return 0;
}

