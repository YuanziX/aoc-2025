package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
)

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func main() {
	file, _ := os.Open("input")
	lineCount, _ := lineCounter(bufio.NewReader(file))
	file.Seek(0, 0)
	sc := bufio.NewScanner(file)

	input := make([][]rune, 0, lineCount)
	for sc.Scan() {
		input = append(input, []rune(sc.Text()))
	}

	log.Println(D4P2(input))
}
