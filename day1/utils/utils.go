package utils

import (
	"bufio"
	"os"
)

func ProcessFile(path *string, lineToValue func(*string) int) int {
	f, e := os.Open(*path)
	if e != nil {
		panic("Cannot open file")
	}

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineVal := lineToValue(&line)
		sum += lineVal
	}

	return sum
}