package main

import (
	"advent/day1/utils"
	"fmt"
)

var values = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func lineToValue(line *string) int {
	lineLen := len(*line)
	var first, last int
loop1:
	for i := 0; i < lineLen; i++ {
		cVal := int((*line)[i])

		if cVal >= 48 && cVal <= 57 {
			first = cVal - 48
			break
		} else {
			for name, val := range values {
				endIdx := i + len(name)
				if endIdx <= lineLen {
					substr := (*line)[i:endIdx]
					if name == substr {
						first = val
						break loop1
					}
				}
			}
		}
	}
loop2:
	for i := lineLen - 1; i >= 0; i-- {
		cVal := int((*line)[i])
		if cVal >= 48 && cVal <= 57 {
			last = cVal - 48
			break
		} else {
			for name, val := range values {
				endIdx := i + len(name)
				if endIdx <= lineLen {
					substr := (*line)[i:endIdx]
					if name == substr {
						last = val
						break loop2
					}
				}
			}
		}
	}

	return first*10 + last
}

func main() {
	path := "input.txt"
	output := utils.ProcessFile(&path, lineToValue)

	fmt.Println(output)
}
