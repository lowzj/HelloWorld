package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/contains-duplicate-iii/
// 220. 存在重复元素 | Contains Duplicate III
// 其他:
//   * 217: https://leetcode-cn.com/problems/contains-duplicate/
//   * 219: https://leetcode-cn.com/problems/contains-duplicate-ii/
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	return containsNearbyAlmostDuplicate0(nums, k, t)
}

//------------------------------------------------------------------------------
// Solution 2
//
// 复杂度分析:
//   * 时间: O(N)
//   * 空间: O(N)
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	return containsNearbyAlmostDuplicate0(nums, k, t)
}

//------------------------------------------------------------------------------
// Solution 1
//
// 类似 Solution 0 的优化思路, 利用滑动窗口和二叉搜索树. 用大小为 k+1 的窗口在数组 A 滑动, 并
// 将窗口内的元素加入到二叉搜索树中, 在新元素 A[i] 入窗口前将 A[i-k-1] 从树中删除, 再查找
// 大于或等于 A[i]-t 的最小元素 x, 若 x <= A[i]+t, 则返回true, 否则将 A[i] 加入到搜索树中.
//
// 复杂度:
//   * 时间: O(N*lgk)
//   * 空间: O(k)
func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	return containsNearbyAlmostDuplicate0(nums, k, t)
}

//------------------------------------------------------------------------------
// Solution 0
//
// 离散化 + 树状数组(Binary Indexed Trees) + 滑动窗口
//
// 数据范围:
//   * 长度: N ∊ [0,2*10^4]
//   * 元素值: A[i] ∊ [-2^31, 2^31-1]
//   * k: [0, 10^4]
//   * t: [0, 2^31-1]
//
// 暴力解法: 对于每一个新入窗口的元素, 计算其与窗口内其他各个元素的差的绝对值.
// 时间: O(N*k), 根据给的数据范围, 最坏是 O(10^8), 超时了.
//
// 优化思路: 用大小为 k+1 的窗口在数组 A 上滑动, 统计窗口内元素在 [A[i]-t,A[i]+t] 的个数;
// 如果大于1, 则返回true. 利用树状数组进行计数, 复杂度 O(lg(MAX(A[0],A[1],...,A[N-1]))).
// 由于元素值范围过大, 直接利用 BIT 申请空间太大. 但是数组长度在 10^4 内, 可以先离散化, 缩小数据
// 范围, 最终计数的时间复杂度: O(lgN).
//
// 步骤:
//   1. 去重并离散化, 得到去重并排序后的数组, 以及离散后的哈希表:
//      * ordered: 去重并升序.
//      * hash: ordered[i] -> i, hash[ordered[i]] = i+1. 加1是为了保证 BIT 可用.
//   2. 构建树状数组(BIT), 用于统计和查询区间内的元素个数:
//      * 范围: [1, len(ordered)]
//      * 大小: n = len(ordered)+1
//      * 操作 add(i,v): 更新区间 [i, n), 增加 v. v 只能是1或-1, 表示入窗口和出窗口.
//      * 操作 sum(i): 查询区间 [1, i] 的元素个数.
//   3. 对于每个新入窗口的元素 A[i]:
//      * A[i] 入窗口: 更新BIT, add(hash[A[i]], 1).
//      * A[i-k-1] 出窗口: 当 i>=k+1 时, add(hash[A[i-k-1]], -1).
//      * 确定查询区间: [A[i]-t, A[i]+t], 映射到离散后的 ordered 下标范围中, 二分查找.
//        * 离散区间(l, r]: 因为 sum(r)-sum(l) 会去掉 l 处的计数.
//        * l: ordered 中小于 A[i]-t 的最大元素的下标 + 1.
//        * r: ordered 中小于或等于 A[i]+t 的最小元素的下标 + 1.
//      * 统计区间内元素个数: sum(r)-sum(l), 表示区间 (l,r], 大于1返回 true.
//
// 复杂度分析:
//   * 时间: O(N*lgN). 去重并离散 O(N*lgN), 窗口滑动 O(N), 计数 O(lgN).
//   * 空间: O(N)
func containsNearbyAlmostDuplicate0(nums []int, k int, t int) bool {
	N := len(nums)
	if N == 0 {
		return false
	}

	// discretize and delete duplicated values
	ordered, hash, n := discretize(nums)
	// Binary Indexed Trees: [1,n]
	NN := n + 1
	tr := make([]int, NN)
	add := func(i, v int) {
		for ; i < NN; i += i & -i {
			tr[i] += v
		}
	}
	sum := func(i int) (sum int) {
		for ; i > 0; i -= i & -i {
			sum += tr[i]
		}
		return sum
	}

	for i, v := range nums {
		add(hash[v], 1)
		if i >= k+1 {
			add(hash[nums[i-k-1]], -1)
		}
		l := sort.SearchInts(ordered, v-t)
		r := sort.SearchInts(ordered, v+t)
		if r < n && ordered[r] == v+t {
			r++
		}
		if sum(r)-sum(l) > 1 {
			return true
		}
	}
	return false
}

func discretize(nums []int) ([]int, map[int]int, int) {
	N := len(nums)
	hash := make(map[int]int)
	// de-dup
	for i := 0; i < N; i++ {
		hash[nums[i]] = i
	}
	ordered, n := make([]int, len(hash)), 0
	for k, _ := range hash {
		ordered[n] = k
		n++
	}
	sort.Ints(ordered)
	for i := 0; i < n; i++ {
		// BIT starts with index 1
		hash[ordered[i]] = i + 1
	}
	return ordered, hash, n
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := []struct {
		nums []int
		k, t int
	}{
		// true
		{
			[]int{1, 2, 3, 1},
			3, 0,
		},
		{
			[]int{1, 0, 1, 1},
			1, 2,
		},
		// false
		{
			nil, 0, 0,
		},
		{
			[]int{1, 2, 1, 1},
			1, 0,
		},
		{
			[]int{1, 5, 9, 1, 5, 9},
			2, 3,
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(containsNearbyAlmostDuplicate0(c.nums, c.k, c.t))
	}
}
