package main

import "strings"

func simplifyPath(path string) string {
	elems := strings.Split(path, "/")
	ans := make([]string, 0, len(elems))
	for _, v := range elems {
		if v == ".." {
			if len(ans) > 0 {
				ans = ans[:len(ans)-1]
			}
			continue
		}
		if v == "." || v == "" {
			continue
		}
		ans = append(ans, v)
	}
	str := ""
	for _, v := range ans {
		str += "/" + v
	}
	if str == "" {
		return "/"
	}
	return str
}
