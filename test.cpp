#include <iostream>
#include <vector>
#include <map>

using namespace std;

#define GENERATE(id, listName, typeName, funName) \
  namespace {                                     \
    const int listName = id;                      \
    A* obj_##funName = new A(id, funName);        \
  }                                               \

class Manager;
class A {
public:
  typedef bool (*CreateFunction)(int id, vector<int>& v);
  A(int id, CreateFunction cf);
  void print() {
    cout << "A print: " << id_ << endl;
  }
  bool create(int id, vector<int>& v) {
    if (createFun_) {
      return createFun_(id, v);
    }
    return false;
  }
  int id() {
    return id_;
  }
protected:
  int id_;
  CreateFunction createFun_;
};

class Manager {
private:
  Manager(){
    cout << "Manager constructor." << endl;
  };
  virtual ~Manager() {
    if (obj_) {
      delete obj_;
    }
    obj_ = NULL;
  }
  static Manager* obj_;
  map<int, A*> fs_;

public:
  static Manager& instance() {
    if (!obj_) {
      obj_ = new Manager;
    }
    return *obj_;
  }

  void add(int id, A* a) {
    if (a != NULL && fs_.insert(make_pair(a->id(), a)).second) {
      cout << "Manager::add id[" << a->id() << "] success!" << endl;
    } else {
      cout << "Manager::add id[" << id << "] failed!" << endl;
    }
  }

  A* get(int id) {
    if (fs_.find(id) != fs_.end()) {
      return fs_.find(id)->second;
    }
    return NULL;
  }

};
Manager* Manager::obj_ = NULL;

A::A(int id, CreateFunction cf): id_(id), createFun_(cf) {
  cout << "A gen: " << id_ << endl;
  Manager::instance().add(id_, this);
}

bool createFriend(int id, vector<int>& v) {
  cout << "createFriend start" << endl;
  v.clear();
  for (int i = 0; i < 10; ++i) {
    v.push_back(i);
  }
  return true;
}
bool createPage(int id, vector<int>& v) {
  cout << "createPage start" << endl;
  v.clear();
  for (int i = 10; i < 20; ++i) {
    v.push_back(i);
  }
  return true;
}

GENERATE(1, LIST_FRIEND, "a", createFriend);
GENERATE(2, LIST_PAGE, "b", createPage);

void print(const vector<int>& v) {
  for (int i = 0; i < v.size(); ++i) {
    cout << v[i] << " ";
  }
  cout << endl;
}
int main(void) {
  A* a = Manager::instance().get(1);
  vector<int> v;
  a->create(1, v);
  print(v);

  a = Manager::instance().get(2);
  a->create(1, v);
  print(v);
  return 0;
}
