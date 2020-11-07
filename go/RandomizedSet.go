package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	dict map[int]int
	vals []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		dict: make(map[int]int),
		vals: make([]int, 0, 1024),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.dict[val]; ok {
		return false
	}
	r.dict[val] = len(r.vals)
	r.vals = append(r.vals, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (r *RandomizedSet) Remove(val int) bool {
	n := len(r.vals) - 1
	i, ok := r.dict[val]
	if !ok || n < 0 {
		return false
	}

	r.vals[i] = r.vals[n]
	r.dict[r.vals[i]] = i

	delete(r.dict, val)
	r.vals = r.vals[:n]
	return ok
}

/** Get a random element from the set. */
func (r *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(r.vals))
	return r.vals[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	cases := [][]int{
		{},
		{},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(c)
	}
}
