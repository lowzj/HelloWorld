#include <iostream>
#include <vector>
#include <set>
#include <unordered_map>

using namespace std;
struct Interval {
    int start;
    int end;
    Interval() : start(0), end(0) {}
    Interval(int s, int e) : start(s), end(e) {}
};
        
class Solution {
#define L(x) ((x)*2+1)
#define R(x) (((x)+1)*2)
#define P(x) (((x)-1)/2)
private:
    struct SegTreeNode {
        int s;
        int e;
        int flag;
        SegTreeNode():s(-1), e(-1), flag(0){}
    };

    SegTreeNode *tree_;
    unordered_map<int, int> m_;
    int *v_;
    int cnt_;

    void build(int start, int end, int pos) {
        if (start > end || pos >= cnt_) return;
        cout << "build --> start:" << start << " end:" << end << " pos:" << pos << endl;
        tree_[pos].s = start;
        tree_[pos].e = end;

        if (start == end) return;
        build(start, (start+end)/2, L(pos));
        build((start+end)/2, end, R(pos));
    }

    void insert(int start, int end, int pos) {
        cout << "insert --> start:" << start << " end:" << end << " pos:" << pos << endl;
        if (start >= end || tree_[pos].flag == 1) return;
        if (start == tree_[pos].s && end == tree_[pos].e) {
            tree_[pos].flag = 1;
            cout << "insert --> ["<<v_[tree_[pos].s]<<","<<v_[tree_[pos].e]<<"]"<<tree_[pos].flag << endl;
            return;
        }
        int l = L(pos);
        int r = R(pos);
        if (l < cnt_ && tree_[l].e > start) {
            insert(start, end, l);
            tree_[pos].flag = tree_[pos].flag || tree_[l].flag;
        }
        if (r < cnt_ && tree_[r].s < end) {
            insert(start, end, r);
            tree_[pos].flag = tree_[pos].flag && tree_[r].flag;
        }
        cout << "insert --> ["<<v_[tree_[pos].s]<<","<<v_[tree_[pos].e]<<"]"<<tree_[pos].flag << endl;
    }
    void build(vector<Interval> &intervals, Interval newInterval) {
        set<int> s;

        int size = intervals.size();
        for (int i = 0; i < size; ++i) {
            s.insert(intervals[i].start);
            s.insert(intervals[i].end);
        }
        s.insert(newInterval.start);
        s.insert(newInterval.end);

        cnt_ = 0;
        v_ = new int[s.size()];
        for (set<int>::iterator i = s.begin(); i != s.end(); ++i) {
            m_.insert(make_pair(*i, cnt_));
            v_[cnt_++] = *i;
            cout << "v_["<<(cnt_-1)<<"]:" << (*i) << " ";
        }
        cout << endl;
        cnt_ *= 4;
        tree_ = new SegTreeNode[cnt_];
        build(0, cnt_-1, 0);

        for (int i = 0; i < size; ++i) {
            insert(m_[intervals[i].start], m_[intervals[i].end], 0);
        }
        insert(m_[newInterval.start], m_[newInterval.end], 0);
    }
    void traverse(int pos, vector<Interval> &v) {
        if (tree_[pos].s == -1 || pos >= cnt_) return;
        cout << "traverse --> [" << tree_[pos].s << ", " << tree_[pos].e << "]" << endl;
        if (tree_[pos].flag) {
            v.push_back(Interval(v_[tree_[pos].s], v_[tree_[pos].e]));
            // cout << "push --> [" << v_[tree_[pos].s] << ", " << v_[tree_[pos].e] << "]" << endl;
            return;
        }
        if (L(pos) < cnt_)
            traverse(L(pos), v);
        if (R(pos) < cnt_)
            traverse(R(pos), v);
    }
public:
    vector<Interval> insert(vector<Interval> &intervals, Interval newInterval) {
        build(intervals, newInterval);

        vector<Interval> ret;
        traverse(0, ret);
        return ret;
    }
};


int main(void) {
    vector<Interval> v;
    v.push_back(Interval(6, 7));
    v.push_back(Interval(8, 8));
    v.push_back(Interval(9, 9));

    Solution s;
    vector<Interval> ret = s.insert(v, Interval(14,22));
    for (int i = 0; i < ret.size(); ++i)
        cout << "["<<ret[i].start<<", "<<ret[i].end<<"]" << endl;
    return 0;
}

