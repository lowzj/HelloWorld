#include <iostream>
#include <stack>

using namespace std;

#define MAXINT 0X7FFFFFFF

struct MinStack {
    stack<int> _s;
    int _min;

    MinStack() {
        _min = MAXINT;
    }

    void push(int v) {
        cout << "value: " << v << " _min: " << _min << " push: " << (v-_min);
        _s.push(v - _min);
        if (_min > v) {
            _min = v;
        }
        cout << " after_min: " << _min << endl;
    }

    void pop() {
        cout << "pop_value: " << top() << " _min: " << _min;
        if (_s.top() < 0) {
            _min = _min - _s.top();
        }
        _s.pop();
        cout << " after_min: " << _min << endl;
    }

    int min() {
        cout << "min: " << _min << endl;
        return _min;
    }

    int top() {
        if (_s.top() < 0) {
            return _min;
        }
        return _s.top() + _min;
    }
};

int main(void) {
    MinStack s;
    int vs[] = {100, 300, 90, 91, 92, 93};
    for (int i = 0; i < 6; ++i) {
        s.push(vs[i]);
    }
    s.top();
    s.min();
    for (int i = 0; i < 6; ++i) {
        s.pop();
    }
    
    return 0;
}
