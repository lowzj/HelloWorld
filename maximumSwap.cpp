#include <iostream>

using namespace std;

class Solution {
public:
    int restore(int x[], int len) {
        int ret = 0;
        for (int i = len-1; i >= 0; --i) ret = x[i] + ret * 10;
        return ret;
    }
    int maximumSwap(int num) {
        int x[10] = {0};
        int l = 0;
        int tmp = num;
        for (l = 0; l < 10 & tmp > 0; ++l) {
            x[l] = tmp % 10;
            tmp /= 10;
        }
        for (int i = l-1; i > 0; --i) {
            int max = 0;
            for (int j = i-1; j >= 0; --j) max = x[max] > x[j] ? max : j;
            if (x[max] > x[i]) {
                tmp = x[max];
                x[max] = x[i];
                x[i] = tmp;
                return restore(x, l);
           }
        }
        return num;
    }
};

int main(void) {
    Solution s;
    cout << s.maximumSwap(1993) << endl;
    return 0;
}
