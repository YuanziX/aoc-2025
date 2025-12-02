package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func P1() {
	file, err := os.Open("input")
	if err != nil {
		log.Panicf("could not open file: %v", err)
	}

	sc := bufio.NewScanner(file)
	ctr := 0
	dial := 50

	for sc.Scan() {
		in := sc.Text()
		num, _ := strconv.ParseInt(in[1:], 10, 64)

		if in[0] == 'L' {
			dial = (dial - int(num)) % 100
			if dial < 0 {
				dial += 100
			}
		} else {
			dial = (dial + int(num)) % 100
		}
		if dial == 0 {
			ctr++
		}
	}

	println(ctr)
}
