#include <iostream>
#include <unordered_set>
#include <vector>
#include <string>
#include <stack>

using namespace std;

class Solution {
public:
    int ladderLength(string start, string end, unordered_set<string> &dict) {
        int len = start.length();

        int m[len];
        for (int i = 0; i < len; ++i) m[i] = 0;

        string tmp = start;
        int cnt = 0;
        while (tmp != end) {
            int tmpCnt = cnt;
            for (int i = 0; i < len; ++i) {
                if (!m[i] && tmp[i] != end[i]) {
                    char c = tmp[i];
                    tmp[i] = end[i];
                    cout << "tmp:" << tmp << " cnt:" << cnt << endl;
                    if (dict.find(tmp) != dict.end() || tmp == end) {
                        m[i] = 1;
                        ++cnt;
                        break;
                    }
                    tmp[i] = c;
                }
            }
            cout << "tmp:" << tmp << " cnt:" << cnt << endl;
            if (cnt == tmpCnt)
                return tmp == end ? cnt : -1;
        }
        return tmp == end ? cnt : -1;
    }
};

int main(void) {
    string start("hit");
    string end("cog");
    string ss[] = {"hot","dot","dog","lot","log"};
    unordered_set<string> dict;
    for (int i = 0 ; i < sizeof(ss)/sizeof(string); ++i)
        dict.insert(ss[i]);

    Solution s;
    cout << s.ladderLength(start, end, dict) << endl;
    return 0;
}
