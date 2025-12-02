package main

import (
	"log"
	"strconv"
	"strings"
)

func isValid(num string) bool {
	for i := range len(num) {
		if strings.ReplaceAll(num, num[:i], "") == "" {
			return false
		}
	}
	return true
}

func P2(input string) {
	ranges := strings.Split(input, ",")
	result := 0

	for _, r := range ranges {
		vals := strings.Split(r, "-")
		left, _ := strconv.Atoi(vals[0])
		right, _ := strconv.Atoi(vals[1])

		for ; left <= right; left++ {
			str := strconv.Itoa(left)

			if !isValid(str) {
				result += left
			}
		}
	}

	log.Println(result)
}
