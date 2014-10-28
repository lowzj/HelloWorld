#include <iostream>
#include <string>
#include <vector>
#include <unordered_set>
#include <unordered_map>

using namespace std;

class Solution {
public:
    typedef unordered_map<int, vector<int> > WordMap;
    typedef WordMap::iterator WordMapItr;

    vector<string> assemble(WordMap& m, string& s, int n) {
        vector<string> v;
        if (m.find(n) == m.end()) {
            return v;
        }

        vector<int>& tmp = m[n];
        for (int i = tmp.size()-1; i >= 0; --i) {
            string str = s.substr(tmp[i]+1, n-tmp[i]);
            if (tmp[i] != -1) {
                vector<string> ss = assemble(m, s, tmp[i]);
                for (int j = ss.size()-1; j >= 0; --j) {
                    v.push_back(ss[j]+" "+str);
                }
            } else {
                v.push_back(str);
            }
        }
        return v;
    }
    vector<string> wordBreak(string s, unordered_set<string> &dict) {

        WordMap m;
        int len = s.size();

        for (int i = 0; i < len; ++i) {
            vector<int> v;
            string tmp = s.substr(0, i+1);
            if (dict.end() != dict.find(tmp)) {
                v.push_back(-1);
            }
            for (WordMapItr it = m.begin(); it != m.end(); ++it) {
                int idx = it->first;
                string tmp = s.substr(idx+1, i-idx);
                if (dict.end() != dict.find(tmp)) {
                    v.push_back(idx);
                }
            }
            if (!v.empty())
                m.insert(make_pair(i, v));
        }

        return assemble(m, s, len-1);
    }

    // Memory Limit Exceeded
    vector<string> wordBreak2(string s, unordered_set<string> &dict) {
        int len = s.size();
        vector<vector<string> > segs;
        for (int i = 0; i < len; ++i) {
            vector<string> seg;
            if (dict.end() != dict.find(s.substr(0, i+1))) {
                seg.push_back(s.substr(0, i+1));
            }
            for (int j = segs.size()-1; j >= 0; --j) {
                if (!segs[j].empty()) {
                    string tmp = s.substr(j+1, i-j+1);
                    if (dict.find(tmp) != dict.end()) {
                        for (int k = segs[j].size()-1; k>=0; --k) {
                            seg.push_back(segs[j][k]+" "+tmp);
                        }
                    }
                }
            }
            segs.push_back(seg);
        }
        return segs[segs.size()-1];
    }
};

int main(void) {
    string str = "leetcode";
    unordered_set<string> dict;
    dict.insert("leet");
    dict.insert("code");

    Solution s;
    vector<string> v = s.wordBreak(str, dict);
    for (int i = 0; i < v.size(); ++i) 
        cout << v[i] << endl;
    return 0;
}

