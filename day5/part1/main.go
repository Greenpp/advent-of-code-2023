package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	sc "strconv"
	s "strings"
)

type valueMap struct {
	sourceStart int
	range_      int
	shift       int
}

func parseInts(values []string) []int {
	parsed := []int{}
	for _, v := range values {
		p, e := sc.Atoi(s.TrimSpace(v))
		if e != nil {
			log.Fatalf("Failed to convert %s to int\n", v)
		}
		parsed = append(parsed, p)
	}

	return parsed
}

func mapValue(v int, maps []valueMap) int {
	for _, m := range maps {
		if v >= m.sourceStart && v < (m.sourceStart+m.range_) {
			return v + m.shift
		}
	}
	return v
}

func mapValues(values []int, maps []valueMap) []int {
	mappedValues := []int{}
	for _, v := range values {
		mappedValues = append(mappedValues, mapValue(v, maps))
	}
	return mappedValues
}

func processLines(path *string) int {
	f, e := os.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	seedsLine := scanner.Text()
	values := parseInts(s.Split(seedsLine, " ")[1:])
	maps := []valueMap{}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] < 48 || line[0] > 57 {
			if len(maps) > 0 {
				values = mapValues(values, maps)
				maps = []valueMap{}
			}
		} else {
			mapParams := parseInts(s.Split(line, " "))
			maps = append(maps, valueMap{
				sourceStart: mapParams[1],
				range_:      mapParams[2],
				shift:       mapParams[0] - mapParams[1],
			})
		}
	}
	values = mapValues(values, maps)
	return slices.Min(values)
}

func main() {
	path := "input.txt"
	output := processLines(&path)
	fmt.Println(output)
}
