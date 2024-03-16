package main

func firstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}

	mem := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		val, ok := mem[s[i]]
		if ok {
			if val >= 0 {
				mem[s[i]] = -1
			}
		} else {
			mem[s[i]] = i
		}
	}
	ans := 100001
	for _, v := range mem {
		if v > -1 {
			ans = min(ans, v)
		}
	}
	if ans != 100001 {
		return ans
	}
	return -1
}
