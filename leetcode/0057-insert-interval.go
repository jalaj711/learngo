package main

func findpos(intervals [][]int, new []int) int {
	low := 0
	high := len(intervals) - 1
	var mid int
	for low < high {
		mid = (low + high) / 2
		if mid == high {
			return mid
		}
		if intervals[mid][0] <= new[0] {
			if intervals[mid+1][0] >= new[0] {
				return mid + 1
			}
			low = mid + 1
		} else {
			high = mid
		}
	}
	if intervals[high][0] > new[0] {
		return high
	}
	return high + 1
}

func insert(intervals [][]int, newInterval []int) [][]int {

	if len(intervals) == 0 {
		return [][]int{{newInterval[0], newInterval[1]}}
	}

	pos := findpos(intervals, newInterval)
	inserted := make([][]int, 0, len(intervals)+1)
	inserted = append(inserted, intervals[:pos]...)
	if pos > 0 && inserted[pos-1][1] >= newInterval[0] {
		inserted[pos-1][1] = max(inserted[pos-1][1], newInterval[1])
	} else {
		inserted = append(inserted, newInterval)
	}
	for pos < len(intervals) && intervals[pos][0] <= inserted[len(inserted)-1][1] {
		inserted[len(inserted)-1][1] = max(inserted[len(inserted)-1][1], intervals[pos][1])
		pos++
	}
	if pos < len(intervals) {
		inserted = append(inserted, intervals[pos:]...)

	}
	return inserted
}

// func main() {

// 	insert([][]int{{2, 6}, {7, 9}}, []int{15, 18})
// 	insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
// 	insert([][]int{{0, 3}, {4, 5}, {6, 9}}, []int{2, 5})
// 	insert([][]int{{0, 2}, {3, 4}, {4, 5}, {6, 9}}, []int{2, 3})
// 	insert([][]int{{1, 3}}, []int{4, 5})
// 	insert([][]int{{6, 9}}, []int{11, 15})
// }
