#include <iostream>

using namespace std;

int maxProduct(int A[], int n) {
    if (!A) return 0;
    int max = A[0];
    int p1, p2, flag;
    p1 = p2 = flag = 1;
    for (int i = 0; i < n; ++i) {
        p1 *= A[i];
        if (!flag) {
            p2 *= A[i];
        }
        if (max < p1)
            max = p1;
        if (!flag && max < p2)
            max = p2;

        if (A[i] < 0 && flag) {
            flag = 0;
        }
        if (A[i] == 0) {
            p1 = p2 = flag = 1;
        }
    }
    return max;
}

int t(int A[], int n) {
}

int main(void) {
    int A[] = {2, 3, -4, 5, 0, -6, -3};
    cout << maxProduct(A, 7) << endl;
    return 0;
}
