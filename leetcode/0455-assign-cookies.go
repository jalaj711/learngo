package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	gp := 0
	sp := 0

	for gp < len(g) && sp < len(s) {
		if g[gp] <= s[sp] {
			gp++
		}
		sp++
	}

	return gp
}
