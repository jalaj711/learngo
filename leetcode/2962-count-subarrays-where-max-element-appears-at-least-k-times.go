package main

import "slices"

func countSubarrays(nums []int, k int) int64 {
	m := slices.Max(nums)
	var l, r, c, a int64
	n := int64(len(nums))
	_k := int64(k)
	for r < n {
		if nums[r] == m {
			c++
			for c == _k && l <= r {
				a += n - r
				if nums[l] == m {
					c--
				}
				l++
			}
		}
		r++
	}
	return a
}

//func main() {
//	countSubarrays([]int{2, 4, 6, 5, 6}, 1)
//	countSubarrays([]int{1, 3, 2, 3, 3}, 2)
//}
