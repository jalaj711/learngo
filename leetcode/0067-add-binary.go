package main

import "fmt"

func addBinary(a string, b string) string {
	ans := ""
	a_sz := len(a)
	b_sz := len(b)
	carry := 0
	ptr := 0
	sum := 0

	for ptr < a_sz && ptr < b_sz {
		sum = int(a[a_sz-ptr-1]-'0') + int(b[b_sz-ptr-1]-'0') + carry
		ans = fmt.Sprintf("%d", sum%2) + ans
		carry = sum / 2
		ptr++
	}
	for ptr < a_sz {
		sum = int(a[a_sz-ptr-1]-'0') + carry
		ans = fmt.Sprintf("%d", sum%2) + ans
		carry = sum / 2
		ptr++
	}
	for ptr < b_sz {
		sum = int(b[b_sz-ptr-1]-'0') + carry
		ans = fmt.Sprintf("%d", sum%2) + ans
		carry = sum / 2
		ptr++
	}
	if carry != 0 {
		ans = "1" + ans
	}
	return ans
}
