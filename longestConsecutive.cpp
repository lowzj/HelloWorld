#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
public:
    int longestConsecutive(vector<int> &num) {
        unordered_map<int, int> m;
        int len = num.size();

        int v[len];
        int mark[len];

        for (int i = 0; i < len; ++i) {
            m.insert(make_pair(num[i], i));
            v[i] = -1;
            mark[i] = 0;
        }
        for (int i = 0; i < len; ++i) {
            if (m.find(num[i]-1) != m.end()) {
                int smallIdx = m[num[i]-1];
                v[i] = smallIdx;
            }
        }

        int maxSeq = 1;
        for (int i = 0; i < len; ++i) {
            if (mark[i]) continue;
            mark[i] = 1;

            int p = v[i];
            while (p != -1) {
                if (mark[p]) {
                    mark[i] += mark[p];
                    break;
                }
                else {
                    ++mark[i];
                    mark[p] = 1;
                }
                p = v[p];
            }
            maxSeq = mark[i] > maxSeq ? mark[i] : maxSeq;
        }
        return maxSeq;
    }
};

int main(void) {
    int a[] = {1,2,0,1};
    vector<int> v;
    for (int i = 0; i < sizeof(a)/sizeof(int); ++i)
        v.push_back(a[i]);

    Solution s;
    cout << s.longestConsecutive(v) << endl;

    return 0;
}

