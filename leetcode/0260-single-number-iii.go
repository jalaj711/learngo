package main

func singleNumber(nums []int) []int {
	if len(nums) == 2 {
		return nums
	}
	x := nums[0]
	for i := 1; i < len(nums); i++ {
		x ^= nums[i]
	}
	c := 1
	for (x & 1) != 1 {
		x >>= 1
		c <<= 1
	}
	ans := make([]int, 2)
	f1 := false
	f2 := false
	for j, _ := range nums {
		if (nums[j] & c) != 0 {
			if f1 {
				ans[0] ^= nums[j]
			} else {
				f1 = true
				ans[0] = nums[j]
			}
		} else {
			if f2 {
				ans[1] ^= nums[j]
			} else {
				f2 = true
				ans[1] = nums[j]
			}
		}
	}
	return ans
}
