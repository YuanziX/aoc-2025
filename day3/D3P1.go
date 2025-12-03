package main

import (
	"strconv"
)

func D3P1(input string) int {
	var index int
	first := -1

	for i := range len(input) {
		n, _ := strconv.Atoi(string(input[i]))
		if n > first && i != len(input)-1 {
			index = i
			first = n
		}
	}

	second := -1

	for j := range len(input) - index - 1 {
		n, _ := strconv.Atoi(string(input[j+index+1]))
		if n > second {
			second = n
		}
	}

	return 10*first + second
}
