package main

//func __recur(i, n int, nums, dp []int) {
//	if i >= n-1 {
//		return
//	}
//	dp[i] = 0
//	for j := 0; j <= nums[i]; j++ {
//		if i+j >= n-1 {
//			dp[i] = 1
//			break
//		}
//		if dp[i+j] < 0 {
//			__recur(i+j, n, nums, dp)
//		}
//		if dp[i+j] > 0 {
//			if dp[i] == 0 || dp[i+j]+1 < dp[i] {
//				dp[i] = 1 + dp[i+j]
//			}
//		}
//	}
//}
//
//func jump(nums []int) int {
//	n := len(nums)
//	if n == 1 {
//		return 0
//	}
//	dp := make([]int, n)
//	for i, _ := range dp {
//		dp[i] = -1
//	}
//	__recur(0, n, nums, dp)
//	return dp[0]
//}

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	nums[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		j := nums[i]
		nums[i] = 100000
		for ; j > 0; j-- {
			if i+j >= n-1 {
				nums[i] = 1
				break
			} else {
				nums[i] = min(nums[i], 1+nums[i+j])
			}
		}
	}
	return nums[0]
}
