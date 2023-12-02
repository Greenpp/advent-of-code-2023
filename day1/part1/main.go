package main

import (
	"advent/day1/utils"
	"fmt"
)

func lineToValue(line *string) int {
	lineLen := len(*line)
	var first, last int
	for i := 0; i < lineLen; i++ {
		cVal := int((*line)[i])
		if cVal >= 48 && cVal <= 57 {
			first = cVal - 48
			break
		}
	}
	for i := lineLen - 1; i >= 0; i-- {
		cVal := int((*line)[i])
		if cVal >= 48 && cVal <= 57 {
			last = cVal - 48
			break
		}
	}

	return first*10 + last
}

func main() {
	path := "input.txt"
	sum := utils.ProcessFile(&path, lineToValue)

	fmt.Println(sum)
}
