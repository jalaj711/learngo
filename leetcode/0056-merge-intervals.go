package main

import "sort"

type SortIntervals [][]int

func (a SortIntervals) Len() int           { return len(a) }
func (a SortIntervals) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortIntervals) Less(i, j int) bool { return a[i][0] < a[j][0] }

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Sort(SortIntervals(intervals))
	ans := make([][]int, 1, len(intervals))
	ans[0] = intervals[0]
	ptr := 0
	for _, v := range intervals {
		if ans[ptr][1] >= v[0] {
			ans[ptr][1] = max(v[1], ans[ptr][1])
		} else {
			ans = append(ans, v)
			ptr++
		}
	}
	return ans
}
