#include <iostream>
#include <string>

using namespace std;

void reverse(string& str, int s, int e) {
    if (s < 0 || e >= str.length()) {
        return;
    }
    while (s < e) {
        char tmp = str[s];
        str[s] = str[e];
        str[e] = tmp;
        ++s;
        --e;
    }
}
void reverseWords(string& s) {
    int i, j;
    int len = s.length();

    string tmp = "";
    i = j = 0;
    while (j < len) {
        while (j < len && s[j] == ' ') ++j;
        if (j < len && tmp.length() > 0) {
            tmp += " ";
        }
        while (j < len && s[j] != ' ') {
            tmp += s[j++];
        }
    }
    s = tmp;
    len = s.length();
    reverse(s, 0, len - 1);

    while (j <= len) {
        if (j == len || s[j] == ' ') {
            reverse(s, i, j-1);
            i = j = j + 1;
        }
        ++j;
    }
}

int main(void) {
    string s = " bc   da  ";
    reverseWords(s);
    cout << s << "||" << endl;
    return 0;
}
