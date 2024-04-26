package main

func maxSubArray(nums []int) int {
	n := len(nums)
	ans := nums[n-1]
	for j := n - 2; j >= 0; j-- {
		if nums[j+1] > 0 {
			nums[j] += nums[j+1]
		}
		ans = max(ans, nums[j])
	}
	return ans
}
