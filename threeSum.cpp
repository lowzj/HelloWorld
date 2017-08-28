#include <iostream>
#include <vector>
#include <set>
#include <algorithm>

using namespace std;

// https://leetcode.com/problems/3sum/description
class Solution {
public:
    int biSearch(vector<int>& v, int start, int value) {
        int low = start;
        int high = v.size() - 1;
        int mid;
        while (low <= high) {
            mid = (low + high) / 2;
            if (v[mid] == value) return mid;
            if (v[mid] > value) {
                high = mid - 1;
            } else low = mid + 1;
        }
        return -1;
    }

    vector<vector<int> > threeSum(vector<int>& vv) {
        vector<vector<int> > res;
        sort(vv.begin(), vv.end());
        vector<int> v;
        int count = 0;
        for (int i = 0; i < vv.size(); ++i) {
            if (count == 3 && vv[i] == vv[i-1]) continue;
            v.push_back(vv[i]);
            if (i > 0 && vv[i] != vv[i-1]) count = 0;
            ++count;
        }

        if (v.size() < 3) return res;
        int max = v[v.size() - 1];
        set<long> mark;

        for (int i = 0; i < v.size() - 1; ++i) {
            for (int j = i + 1; j < v.size(); ++j) {
                int itr = biSearch(v, j + 1, 0-v[i]-v[j]);
                if (itr != -1) {
                    long key = v[i] * max + v[j];
                    if (mark.find(key) == mark.end()) {
                        res.push_back(vector<int>{v[i], v[j], v[itr]});
                        mark.insert(key);
                    }
                }
            }
        }
        return res;
    }
};

int main(void) {
    vector<int> v{-1, 0, 1, 2, -1, -4};
    Solution s;
    vector<vector<int> > vvv = s.threeSum(v);
    for (int i = 0; i < vvv.size(); ++i) {
        for (int j = 0; j < vvv[i].size(); ++j)
            cout << vvv[i][j] << " ";
        cout << endl;
    }
    return 0;
}
