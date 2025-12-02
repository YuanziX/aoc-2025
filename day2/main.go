package main

import (
	"os"
)

func main() {
	file, _ := os.ReadFile("input")
	P2(string(file))
}
