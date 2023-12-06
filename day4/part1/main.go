package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	sc "strconv"
	s "strings"
)

func parseInts(values *string) *map[int]bool {
	splitValues := s.Split(s.TrimSpace(*values), " ")
	parsed := map[int]bool{}
	for _, valTxt := range splitValues {
		if len(valTxt) == 0 {
			continue
		}
		v, e := sc.Atoi(valTxt)
		if e != nil {
			log.Fatalf("Failed to parse '%s' to int\n", valTxt)
		}
		parsed[v] = true
	}

	return &parsed
}

func processLine(line *string) int {
	allValuesTxt := s.Split(*line, ":")[1]
	valuesTxt := s.Split(allValuesTxt, "|")

	winningValues := parseInts(&valuesTxt[0])
	foundValues := parseInts(&valuesTxt[1])

	shared := map[int]bool{}
	var s1, s2 *map[int]bool
	if len(*foundValues) > len(*winningValues) {
		s1, s2 = winningValues, foundValues
	} else {
		s1, s2 = foundValues, winningValues
	}
	for v, _ := range *s1 {
		if (*s2)[v] {
			shared[v] = true
		}
	}

	if len(shared) == 0 {
		return 0
	}
	return int(math.Pow(2, float64(len(shared)-1)))
}

func processLines(path *string) int {
	f, e := os.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += processLine(&line)
	}

	return sum
}

func main() {
	path := "input.txt"
	output := processLines(&path)
	fmt.Println(output)
}
