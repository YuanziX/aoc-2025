package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, _ := os.Open("input")
	sc := bufio.NewScanner(file)

	sum := 0
	for sc.Scan() {
		sum += D3P2(sc.Text())
	}

	log.Println(sum)
}
