package main

import (
	"log"
	"strconv"
	"strings"
)

func P1(input string) {
	ranges := strings.Split(input, ",")
	result := 0

	for _, r := range ranges {
		vals := strings.Split(r, "-")
		left, _ := strconv.Atoi(vals[0])
		right, _ := strconv.Atoi(vals[1])

		for ; left <= right; left++ {
			str := strconv.Itoa(left)

			if len(str)%2 != 0 {
				continue
			}

			half := len(str) / 2

			if str[:half] == str[half:] {
				result += left
			}
		}
	}

	log.Println(result)
}
