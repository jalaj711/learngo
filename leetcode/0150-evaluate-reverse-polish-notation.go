package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _, v := range tokens {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			op1 := stack[len(stack)-2]
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch v {
			case "+":
				stack = append(stack, op1+op2)
			case "-":
				stack = append(stack, op1-op2)
			case "*":
				stack = append(stack, op1*op2)
			case "/":
				stack = append(stack, op1/op2)
			}
		} else {
			val, _ := strconv.Atoi(v)
			stack = append(stack, val)
		}

	}
	return stack[0]
}

// func main() {
// 	evalRPN([]string{"4", "13", "5", "/", "+"})
// }
