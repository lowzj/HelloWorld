#include <algorithm>
#include <iostream>
#include <queue>
#include <stack>
#include <vector>
#include <map>

using namespace std;

// https://leetcode.com/problems/search-in-rotated-sorted-array/description
class Solution {
public:
    int search(vector<int>& nums, int target) {
        int l = 0;
        int h = nums.size() - 1;
        while (l <= h) {
            int m = (l + h) / 2;
            if (target == nums[m]) return m;
            else if (nums[l] > nums[h]) {
                if (target >= nums[l]) {
                    if (target < nums[m] || nums[m] < nums[h]) h = m - 1;
                    else l = m + 1;
                } else {
                    if (target > nums[m] || nums[m] > nums[h]) l = m + 1;
                    else h = m - 1;
                }
            } else {
                if (target > nums[m]) l = m + 1;
                else h = m - 1;
            }
        }
        return -1;
    }
};

int main(void) {
    Solution s;
    vector<int> nums{4,5,6,7,0,1,2};
    cout << 0 << " " << s.search(nums, 0) << endl;
    cout << 1 << " " << s.search(nums, 1) << endl;
    cout << 2 << " " << s.search(nums, 2) << endl;
    cout << 3 << " " << s.search(nums, 3) << endl;
    cout << 4 << " " << s.search(nums, 4) << endl;
    cout << 5 << " " << s.search(nums, 5) << endl;
    cout << 6 << " " << s.search(nums, 6) << endl;
    cout << 7 << " " << s.search(nums, 7) << endl;
    cout << 8 << " " << s.search(nums, 8) << endl;
    return 0;
}
