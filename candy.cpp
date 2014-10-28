#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;

class Solution {
public:
    struct Child {
        int idx;
        int rating;
        int candy;
        Child(int i, int r, int c):idx(i),rating(r),candy(c) {}
        bool operator<(const Child& c) const {
            return rating < c.rating;
        }
    };

    int candy(vector<int> &ratings) {
        vector<Child> vSort;
        vector<Child> v;
        int len = ratings.size();

        for (int i = 0; i < len; ++i) {
            v.push_back(Child(i, ratings[i], 1));
            vSort.push_back(Child(i, ratings[i], 0));
        }

        sort(vSort.begin(), vSort.end());
        int candy = 0;
        for (int i = 0; i < len; ++i) {
            Child& c = v[vSort[i].idx];
            if (c.idx-1 >= 0 && v[c.idx-1].rating < c.rating) {
                c.candy = v[c.idx-1].candy + 1;
            }
            if (c.idx+1 < len && v[c.idx+1].rating < c.rating) {
                c.candy = c.candy > v[c.idx+1].candy ? c.candy : v[c.idx+1].candy + 1;
            }
            candy += c.candy;
        }
        return candy;
    }
};

int main(void) {
    vector<int> v;
    v.push_back(1);
    v.push_back(2);
    v.push_back(3);
    Solution s;
    cout << s.candy(v) << endl;
    return 0;
}

