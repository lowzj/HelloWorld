#include <iostream>
#include <vector>
#include <stack>

using namespace std;

class Solution {
public:
    void mark(int r, int c, int row, int col, vector<vector<char> > &b) {
        if (r < 0 || r >= row || c < 0 || c >= col 
                || b[r][c] == 'X' || b[r][c] == '1') return;
        stack<pair<int, int> > s;
        s.push(make_pair(r, c));
        while (!s.empty()) {
            pair<int,int> rc = s.top();
            s.pop();
            r = rc.first;
            c = rc.second;
            b[r][c] = '1';
            if (r+1 < row && b[r+1][c] != 'X' && b[r+1][c] != '1')
                s.push(make_pair(r+1, c));
            if (r-1 >= 0 && b[r-1][c] != 'X' && b[r-1][c] != '1')
                s.push(make_pair(r-1, c));
            if (c+1 < col && b[r][c+1] != 'X' && b[r][c+1] != '1')
                s.push(make_pair(r, c+1));
            if (c-1 >= 0 && b[r][c-1] != 'X' && b[r][c-1] != '1')
                s.push(make_pair(r, c-1));
        }
    }

    void solve(vector<vector<char> > &board) {
        int row = board.size();
        int col = row ? board[0].size() : 0;

        for (int i = 0; i < row; ++i) {
            if (board[i][0] == 'O')
                mark(i, 0, row, col ,board);
        }
        
        for (int i = 0; i < row; ++i) {
            if (board[i][col-1] == 'O')
                mark(i, col-1, row, col ,board);
        }
        for (int i = 0; i < col; ++i) {
            if (board[0][i] == 'O')
                mark(0, i, row, col ,board);
        }
        for (int i = 0; i < col; ++i) {
            if (board[row-1][i] == 'O')
                mark(row-1, i, row, col ,board);
        }

        for (int i = 0; i < row; ++i) {
            for (int j = 0; j < col; ++j) {
                board[i][j] = board[i][j]=='1' ? 'O' : 'X';
                cout << board[i][j];
            }
            cout << endl;
        }
    }
};

int main(void) {
    string s[] = {"XOOXXXOXOO","XOXXXXXXXX","XXXXOXXXXX","XOXXXOXXXO","OXXXOXOXOX","XXOXXOOXXX","OXXOOXOXXO","OXXXXXOXXX","XOOXXOXXOO","XXXOOXOXXO"};
    vector<vector<char> > board;
    for (int i = 0; i < sizeof(s)/sizeof(string); ++i) {
        vector<char> v;
        for (int j = 0; j < s[i].size(); ++j)
            v.push_back(s[i][j]);
        board.push_back(v);
        cout << s[i] << endl;
    }
    cout << "=======================" << endl;
    Solution so;
    so.solve(board);
    return 0;
}
