package main

import (
	"strconv"
)

func D3P2(input string) int {
	n := len(input)
	k := 12
	stack := make([]int, 0, k)

	for i, ch := range input {
		num, _ := strconv.Atoi(string(ch))

		for len(stack) > 0 && stack[len(stack)-1] < num && len(stack)-1+(n-i) >= k {
			stack = stack[:len(stack)-1]
		}

		if len(stack) < k {
			stack = append(stack, num)
		}
	}

	num := 0
	for _, i := range stack {
		num = num*10 + i
	}
	return num
}
