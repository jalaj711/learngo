package main

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, str := range strs {
		sorted := SortString(str)
		val, ok := anagrams[sorted]
		if ok {
			val = append(val, str)
			anagrams[sorted] = val
		} else {
			anagrams[sorted] = []string{str}
		}

	}
	ans := make([][]string, 0, len(anagrams))
	for _, v := range anagrams {
		ans = append(ans, v)
	}
	return ans
}

// func main() {
// 	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
// }
