package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func D5P2() {
	file, _ := os.Open("input")
	sc := bufio.NewScanner(file)

	ranges := make([]interval, 0, 100)
	change := false

	for sc.Scan(); !change; sc.Scan() {
		inp := sc.Text()
		change = inp == ""
		if change {
			continue
		}

		spl := strings.Split(inp, "-")
		left, _ := strconv.Atoi(spl[0])
		right, _ := strconv.Atoi(spl[1])

		ranges = append(ranges, interval{left: left, right: right})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].left < ranges[j].left
	})

	// merge ranges
	merged := make([]interval, 0, 50)
	merged = append(merged, ranges[0])
	mergedPtr := 0
	for i := range ranges {
		if ranges[i].left <= merged[mergedPtr].right {
			merged[mergedPtr] = interval{
				left: merged[mergedPtr].left, right: max(ranges[i].right, merged[mergedPtr].right),
			}
		} else {
			mergedPtr++
			merged = append(merged, ranges[i])
		}
	}

	count := 0
	for _, rng := range merged {
		count += rng.right - rng.left + 1
	}

	log.Println(count)
}
