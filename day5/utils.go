package main

type interval struct {
	left  int
	right int
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
