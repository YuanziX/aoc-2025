package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func helper(dial, step, dir int) (int, int) {
	ctr := 0

	for range step {
		dial += dir
		if dial == -1 {
			dial = 99
		}

		if dial == 100 {
			dial = 0
		}

		if dial == 0 {
			ctr++
		}
	}

	return dial, ctr
}

func P2() {
	file, err := os.Open("input")
	if err != nil {
		log.Panicf("could not open file: %v", err)
	}

	sc := bufio.NewScanner(file)
	ctr := 0
	dial := 50

	for sc.Scan() {
		in := sc.Text()
		step, _ := strconv.Atoi(in[1:])
		var dir int
		if in[0] == 'L' {
			dir = -1
		} else {
			dir = 1
		}

		d, c := helper(dial, step, dir)
		dial = d
		ctr += c
	}

	println(ctr)
}
