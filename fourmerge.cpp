#include <iostream>
using namespace std;

void print(int* a, int n) {
    for (int i = 0; i < n; ++i) 
        cout << a[i] << " ";
    cout << endl;
}
int zu(int v) {
    return v % 4;
}

void swap(int* a, int* b) {
    int temp = *a;
    *a = *b;
    *b = temp;
}
void fourmeger(int* a, int n) {
    int p = 0; // 0
    int q = 0; // 1
    int x = 0; // 2
    int y = n; // 3
    while (x < y) {
        int fenzu = zu(a[x]);
        if (fenzu == 0) {
            swap(a+x, a+p);
            swap(a+x, a+q);
            ++p, ++q, ++x;
        } else if (fenzu == 1) {
            swap(a+x, a+q);
            ++q, ++x;
        } else if (fenzu == 2) {
            ++x;
        } else {
            --y;
            swap(a+x, a+y);
        }
    }
}

int main(void) {
    int n = 21;
    int* a = new int[n];
    for (int i = 0; i < n; ++i) {
        a[i] = i;
    }
    fourmeger(a, n);
    print(a, n);
    return 0;
}

