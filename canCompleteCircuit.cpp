#include <iostream>
#include <vector>

using namespace std;
class Solution {
public:
    int canCompleteCircuit(vector<int> &gas, vector<int> &cost) {
        int len = gas.size();
        
        int bw = len; // backward
        int fw = 0; // forward
        int cnt = len;
        int gasSum = 0;
        int costSum = 0;

        while (cnt--) {
            if (gasSum+gas[fw] >= costSum+cost[fw]) {
                costSum += cost[fw];
                gasSum += gas[fw];
                ++fw;
            } else {
                --bw;
                costSum += cost[bw];
                gasSum += gas[bw];
            }
            cout << "gas:" << gasSum << " cost:" << costSum << endl;
        }
        return gasSum >= costSum ? (bw%len) : -1;
    }
};

int main(void) {
    vector<int> gas;
    vector<int> cost;
    gas.push_back(4);
    cost.push_back(5);

    Solution s;
    cout << s.canCompleteCircuit(gas, cost) << endl;
    return 0;
}
