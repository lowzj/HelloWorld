#include <iostream>
#include <string>
using namespace std;

extern "C" {
  int print(const string& msg);
};

int print(const string& msg) {
  static int cnt = 0;
  cout << cnt << ": " << msg << endl;
  ++cnt;
}
