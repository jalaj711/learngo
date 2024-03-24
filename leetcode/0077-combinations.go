package main

func _recur(curr []int, ind int, ll int, ul int, ans *[][]int) {
	if ind >= len(curr) || ll > ul {
		return
	}
	if ll == ul {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		tmp[ind] = ll
		*ans = append(*ans, tmp)
		return
	}

	for i := ll; i <= ul-len(curr)+ind+1; i++ {
		curr[ind] = i
		if ind == len(curr)-1 {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			*ans = append(*ans, tmp)
		} else {
			_recur(curr, ind+1, i+1, ul, ans)
		}
	}
}

func combine(n int, k int) [][]int {
	ans := make([][]int, 0, 1)
	curr := make([]int, k)
	_recur(curr, 0, 1, n, &ans)
	return ans
}
