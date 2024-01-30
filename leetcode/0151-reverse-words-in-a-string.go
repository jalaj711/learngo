package main

import "fmt"

func reverseWords(s string) string {
	words := make([]string, 0, 1)
	cword := ""
	for _, r := range s {
		if r == ' ' {
			if cword == "" {
				continue
			} else {
				words = append(words, cword)
				cword = ""
			}
		} else {
			cword += string(r)
		}
	}
	if cword != "" {

		words = append(words, cword)
	}
	fmt.Println(words)
	ans := ""
	for i := len(words) - 1; i > -1; i-- {
		ans += words[i]
		if words[i] != "" && i != 0 {
			ans += " "
		}

	}
	return ans
}

// func main() {
// 	reverseWords("  hello world  ")
// }
