#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int findMin(vector<int> &num) {
        int s = 0;
        int e = num.size() - 1;

        while (s < e) {
            if (num[s] < num[e])
                return num[s];
            int mid = (s+e)/2;
            if (num[mid] < num[e])
                e = mid;
            else 
                s = mid + 1;
        }
        return num[e];
    }

    int findMin2(vector<int> &num) {
        int s = 0;
        int e = num.size() - 1;

        while (s < e) {
            if (num[s] < num[e])
                return num[s];
            int mid = (s+e)/2;
            if (num[mid] < num[e] || num[mid] < num[s])
                e = mid;
            else if (num[mid] > num[e] || num[mid] > num[s])
                s = mid + 1;
            else {
                ++s;
                --e;
            }
        }
        return num[e];
    }
};

int main(void) {
    // int a[] = {5,6,7,-1,0,1,2};
    int a[] = {10,1,10,10};
    vector<int> v;
    for (int i = 0; i < sizeof(a)/sizeof(int); ++i) 
        v.push_back(a[i]);
    Solution s;
    cout << s.findMin2(v) << endl;
    return 0;
}

