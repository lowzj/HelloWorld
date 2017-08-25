#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    double find(vector<int>& g, vector<int>& l) {
        int low = 0;
        int high = g.size();
        int target = (g.size() + l.size() - 1) / 2;
        while (low < high) {
            int gm = (low + high) / 2;
            int lm = target -  gm;
            if ((g.size() + l.size()) % 2 == 1) {
                if (lm == l.size() && g[gm] >= l[lm-1]) {
                    return g[gm] * 1.0;
                } else if (lm > l.size()) {
                    low = gm + 1;
                    continue;
                } else if (lm < 0) {
                    high = gm - 1;
                    continue;
                }
                if (g[gm] <= l[lm] && (lm == 0 || g[gm] >= l[lm-1])) {
                    cout << "1##" << gm << " " << lm << endl;
                    return g[gm] * 1.0;
                } else if (l[lm] <= g[gm] && (gm == 0 || l[lm] >= g[gm-1]))  {
                    cout << "2##" << gm << " " << lm << endl;
                    return l[lm] * 1.0;
                } else if (g[gm] > l[lm]) high = gm - 1;
                else if (g[gm] < l[lm]) low = gm + 1;
            } else {
            }
        }
        return 0.0;
    }

    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int l1 = nums1.size();
        int l2 = nums2.size();
        if (l1 > l2) return find(nums1, nums2);
        else return find(nums2, nums1);
    }
};

int main(void) {
    vector<int> v1{1, 3, 4, 5};
    vector<int> v2{2};
    Solution s;
    cout << s.findMedianSortedArrays(v1, v2) << endl;
    return 0;
}
