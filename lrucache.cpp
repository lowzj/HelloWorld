#include <iostream>
#include <map>
#include <list>
#include <unordered_map>

using namespace std;

class LRUCache{
private:
    // pair<key, value>
    typedef list<pair<int,int> > LRUList;
    typedef LRUList::iterator LRUListIter;
    typedef unordered_map<int, LRUListIter> LRUMap;
    typedef LRUMap::iterator LRUMapIter;

    int     _capacity;
    LRUList _list;
    LRUMap  _map;

public:
    LRUCache(int capacity):_capacity(capacity) {
        if (_capacity < 1)
            _capacity = 1;
    }

    int get(int key) {
        LRUMapIter mit = adjust(key);
        return mit != _map.end() ? mit->second->second : -1;
    }

    void set(int key, int value) {
        LRUMapIter mit = adjust(key);
        if (mit != _map.end()) {
            mit->second->second = value;
        } else {
            _map.insert(make_pair(key, 
                    _list.insert(_list.begin(), make_pair(key, value))));
        }
        evict();
    }

private:
    LRUMapIter adjust(int key) {
        LRUMapIter mit = _map.find(key);
        if (mit != _map.end()) {
            pair<int, int> kv = *(mit->second);
            _list.erase(mit->second);
            mit->second = _list.insert(_list.begin(), kv);
        }
        return mit;
    }
    void evict() {
        while (_map.size() > _capacity) {
            int key = _list.rbegin()->first;
            LRUMapIter mit = _map.find(key);
            if (mit != _map.end()) {
                _list.erase(mit->second);
                _map.erase(mit);
            }
        }
    }
};

int main(void) {
    LRUCache lru(2048);
    for (int i = 0; i < 10000; ++i) {
        lru.set(i, i*10);
    }

    return 0;
}

