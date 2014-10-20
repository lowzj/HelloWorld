#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

struct Point {
    int x;
    int y;
    Point() : x(0), y(0) {}
    Point(int a, int b) : x(a), y(b) {}
};


class Solution {
public:
    int maxPoints(vector<Point> &points) {
        int MAXINT = 0x7fffffff;
        vector<double> v;
        int size = points.size();
        int max = size?1:0;
        for (int i = 0; i < size; ++i) {
            Point& p1 = points[i];
            v.clear();
            int init = 0;
            for (int j = 0; j < size; ++j) {
                Point& p2 = points[j];
                if (p1.x==p2.x && p1.y==p2.y) {
                    ++init;
                    continue;
                }
                int dy = p2.y - p1.y;
                int dx = p2.x - p1.x;
                if (dx == 0) v.push_back(MAXINT*1.0);
                else v.push_back(dy*1.0/dx);
            }
            sort(v.begin(), v.end());
            if (!v.empty()) ++init;
            int cnt = init;
            for (int i = 1; i < v.size(); ++i) {
                if (v[i-1] == v[i]) {
                    ++cnt;
                } else {
                    if (max < cnt) max = cnt;
                    cnt = init;
                }
            }
            if (max < cnt) max = cnt;
        }
        return max;
    }
};

int main(void) {
    vector<Point> p;
    p.push_back(Point(1, 1));
    p.push_back(Point(3, 2));
    p.push_back(Point(5, 3));

    Solution s;
    cout << s.maxPoints(p) << endl;
    return 0;
}
