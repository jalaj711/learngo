package main

func f(d int) int {
	a := 1
	for i := 1; i <= d; i++ {
		a += i + 1
	}
	return a
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)
	l, r := 0, 0
	a := 0
	p := 1
	for l < n && r < n {
		if nums[l] < k {
			p = nums[l]
			for r = l + 1; p < k && r < n; r++ {
				p *= nums[r]
			}
			if p >= k {
				r--
				p /= nums[r]
			}
			a += f(r - l - 1)
			for l < n && r < n && l != r {
				p /= nums[l]
				l++
				for r < n && p*nums[r] < k {
					a += r - l + 1
					p *= nums[r]
					r++
				}
			}
		}
		l++
	}
	return a
}

func main() {
	numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100)        // 8
	numSubarrayProductLessThanK([]int{5, 2, 3, 4, 3}, 120)      // 13
	numSubarrayProductLessThanK([]int{120, 5, 2, 3, 4, 3}, 120) // 13
	numSubarrayProductLessThanK([]int{1, 1, 1}, 2)              // 6
	numSubarrayProductLessThanK([]int{5, 4, 3, 2}, 0)           // 0
}
