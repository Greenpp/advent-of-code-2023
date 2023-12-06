package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	sc "strconv"
)

type label struct {
	value         int
	startMatchIdx int
	endMatchIdx   int
}

func createLabel(line *string, startIdx *int, endIdx *int) *label {
	// end index is set to value len + 1
	valueStr := (*line)[*startIdx:*endIdx]
	value, e := sc.Atoi(valueStr)
	if e != nil {
		log.Fatalf("Failed to convert %s to int", valueStr)
	}
	newLabel := label{
		value:         value,
		startMatchIdx: max((*startIdx)-1, 0),
		endMatchIdx:   min(*endIdx, len(*line)-1),
	}

	return &newLabel
}

func processLine(line *string) ([]*label, []int) {
	// returns list of label idx and list of parts idx
	labelStart := -1
	var parts []int
	var labels []*label
	for i, c := range *line {
		cVal := int(c)
		if cVal >= 48 && cVal <= 57 {
			if labelStart == -1 {
				labelStart = i
			}
		} else {
			if cVal == 42 {
				parts = append(parts, i)
			}
			if labelStart != -1 {
				labels = append(labels, createLabel(line, &labelStart, &i))
				labelStart = -1
			}
		}
	}
	if labelStart != -1 {
		endIdx := len(*line)
		labels = append(labels, createLabel(line, &labelStart, &endIdx))
	}

	return labels, parts
}

func computeLineSum(currentParts *[]int, previousLabels *[]*label, currentLabels *[]*label, nextLabels *[]*label) int {
	sum := 0
	var values []int
	for _, p := range *currentParts {
		values = make([]int, 0)
		for _, l := range *previousLabels {
			if p >= l.startMatchIdx && p <= l.endMatchIdx {
				values = append(values, l.value)
			}
		}
		for _, l := range *currentLabels {
			if p == l.startMatchIdx || p == l.endMatchIdx {
				values = append(values, l.value)
			}
		}
		for _, l := range *nextLabels {
			if p >= l.startMatchIdx && p <= l.endMatchIdx {
				values = append(values, l.value)
			}
		}

		if len(values) == 2 {
			sum += (values[0] * values[1])
		}
	}

	return sum
}

func processLines(path *string) int {
	f, e := os.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0
	var previousLabels, currentLabels, nextLabels []*label
	var currentParts, nextParts []int
	for scanner.Scan() {
		previousLabels = currentLabels
		currentLabels = nextLabels
		currentParts = nextParts

		line := scanner.Text()
		nextLabels, nextParts = processLine(&line)

		sum += computeLineSum(&currentParts, &previousLabels, &currentLabels, &nextLabels)
	}
	previousLabels = currentLabels
	currentLabels = nextLabels
	currentParts = nextParts
	nextLabels = make([]*label, 0)
	sum += computeLineSum(&currentParts, &previousLabels, &currentLabels, &nextLabels)

	return sum
}

func main() {
	path := "input.txt"
	output := processLines(&path)
	fmt.Println(output)
}
