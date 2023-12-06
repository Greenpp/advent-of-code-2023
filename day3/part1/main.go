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
			if cVal != 46 {
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

func computeLineSum(currentLabels *[]*label, previousParts *[]int, currentParts *[]int, nextParts *[]int) int {
	sum := 0
label_loop:
	for _, l := range *currentLabels {
		for _, p := range *previousParts {
			if p >= l.startMatchIdx && p <= l.endMatchIdx {
				sum += l.value
				continue label_loop
			}
		}
		for _, p := range *currentParts {
			if p == l.startMatchIdx || p == l.endMatchIdx {
				sum += l.value
				continue label_loop
			}
		}
		for _, p := range *nextParts {
			if p >= l.startMatchIdx && p <= l.endMatchIdx {
				sum += l.value
				continue label_loop
			}
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
	var previousParts, currentParts, nextParts []int
	var currentLabels, nextLabels []*label
	for scanner.Scan() {
		previousParts = currentParts
		currentParts = nextParts
		currentLabels = nextLabels

		line := scanner.Text()
		nextLabels, nextParts = processLine(&line)

		sum += computeLineSum(&currentLabels, &previousParts, &currentParts, &nextParts)
	}
	previousParts = currentParts
	currentParts = nextParts
	currentLabels = nextLabels
	nextParts = make([]int, 0)
	sum += computeLineSum(&currentLabels, &previousParts, &currentParts, &nextParts)

	return sum
}

func main() {
	path := "input.txt"
	output := processLines(&path)
	fmt.Println(output)
}
