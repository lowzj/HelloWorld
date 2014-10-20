#include <iostream>
#include <stack>
#include <vector>
#include <string>

using namespace std;

bool isOperator(string& s) {
    return (s == "+" || s == "-" ||
            s == "*" || s == "/");
}

int doOperator(string& op, int a, int b) {
    if (op == "+") return a + b;
    if (op == "-") return a - b;
    if (op == "*") return a * b;
    if (op == "/" && b != 0) return a / b;
    return 0;
}

int evalRPN(vector<string>& tokens) {
    stack<int> s;
    for (vector<string>::iterator i = tokens.begin(); i < tokens.end(); ++i) {
        if (isOperator(*i)) {
            int b = s.top();
            s.pop();
            int a = s.top();
            s.pop();
            s.push(doOperator(*i, a, b));
        } else {
            s.push(atoi(i->c_str()));
        }
    }
    return s.top();
}

int main(void) {
    vector<string> tokens;
    tokens.push_back("2");
    tokens.push_back("3");
    tokens.push_back("+");
    tokens.push_back("2");
    tokens.push_back("/");
    cout << evalRPN(tokens) << endl;
    return 0;
}
