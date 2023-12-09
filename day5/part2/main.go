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
	start int
	size  int
	shift int
}

type valueRange struct {
	start int
	size  int
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

func loadSeeds(values []int) []valueRange {
	ranges := []valueRange{}
	for i := 0; i < len(values); i += 2 {
		ranges = append(ranges, valueRange{
			start: values[i],
			size:  values[i+1],
		})
	}
	return ranges
}

func resolveMapping(r valueRange, m valueMap) ([]valueRange, []valueRange) {
	// returns mapped and unmapped ranges
	mapped := []valueRange{}
	if r.start >= (m.start + m.size) {
		// range starts after map
		return mapped, []valueRange{r}
	}
	if (r.start + r.size - 1) < m.start {
		// range ends before map
		return mapped, []valueRange{r}
	}
	unmapped := []valueRange{}
	var sharedStart int
	if r.start < m.start {
		unmapped = append(unmapped, valueRange{
			start: r.start,
			size:  m.start - r.start,
		})
		sharedStart = m.start
	} else {
		sharedStart = r.start
	}
	var sharedEnd int
	if (r.start + r.size) > (m.start + m.size) {
		unmapped = append(unmapped, valueRange{
			start: m.start + m.size,
			size:  (r.start + r.size) - (m.start + m.size),
		})
		sharedEnd = m.start + m.size
	} else {
		sharedEnd = r.start + r.size
	}

	mapped = append(mapped, valueRange{
		start: sharedStart + m.shift,
		size:  sharedEnd - sharedStart,
	})
	return mapped, unmapped
}

func mapRange(r valueRange, maps []valueMap) []valueRange {
	unmapped := []valueRange{r}
	mapped := []valueRange{}
	for _, m := range maps {
		if len(unmapped) == 0 {
			return mapped
		}
		allMapUnmapped := []valueRange{}
		for len(unmapped) > 0 {
			um := unmapped[0]
			unmapped = unmapped[1:]
			mapMapped, mapUnmapped := resolveMapping(um, m)
			mapped = append(mapped, mapMapped...)
			allMapUnmapped = append(allMapUnmapped, mapUnmapped...)
		}
		unmapped = allMapUnmapped
	}
	return append(mapped, unmapped...)
}

func mapRanges(ranges []valueRange, maps []valueMap) []valueRange {
	mapped := []valueRange{}
	for _, r := range ranges {
		mapped = append(mapped, mapRange(r, maps)...)
	}
	return mapped
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
	seedsValues := parseInts(s.Split(seedsLine, " ")[1:])
	values := loadSeeds(seedsValues)
	maps := []valueMap{}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] < 48 || line[0] > 57 {
			if len(maps) > 0 {
				values = mapRanges(values, maps)
				maps = []valueMap{}
			}
		} else {
			mapParams := parseInts(s.Split(line, " "))
			maps = append(maps, valueMap{
				start: mapParams[1],
				size:  mapParams[2],
				shift: mapParams[0] - mapParams[1],
			})
		}
	}
	values = mapRanges(values, maps)
	startValues := []int{}
	for _, r := range values {
		startValues = append(startValues, r.start)
	}
	return slices.Min(startValues)
}

func main() {
	path := "input.txt"
	output := processLines(&path)
	fmt.Println(output)
}
