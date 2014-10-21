#include <iostream>
#include <algorithm>

using namespace std;

class Solution {
public:
    int singleNumber2(int A[], int n) {
        sort(A, A+n);
        int cnt = 1;
        for (int i=1; i < n; ++i) {
            if (A[i] == A[i-1]) ++cnt;
            else if (cnt == 1) return A[i-1];
            else cnt = 1;
        }
        return A[n-1];
    }

    // corndag writed
    int singleNumber(int A[], int n) {
        if(n == 1) return A[0];
        // A[0] is correct to start
        // Take care of processing A[1]
        A[0] ^= A[1];
        // Set A[1] to either 0 or itself
        A[1] = (A[0]^A[1])&A[1];

        // Continue with algorithm as normal
        for(int i = 2; i < n; ++i){
            A[1] |= A[0]&A[i];
            A[0] ^= A[i];
            A[2] = ~(A[0]&A[1]);
            A[0] &= A[2];
            A[1] &= A[2];
        }
        return A[0];
    }

};

int main(void) {
    return 0;
}

