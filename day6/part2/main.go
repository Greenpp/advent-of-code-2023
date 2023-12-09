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

func parseInts(values []string) []int {
	parsed := []int{}
	for _, v := range values {
		if len(v) == 0 {
			continue
		}
		p, e := sc.Atoi(v)
		if e != nil {
			log.Fatalf("Failed to convert %s to int\n", v)
		}
		parsed = append(parsed, p)
	}
	return parsed
}

func getValues(scanner *bufio.Scanner) []int {
	scanner.Scan()
	line := scanner.Text()
	merged := s.ReplaceAll(line, " ", "")
	values := s.Split(merged, ":")
	return parseInts(values[1:])
}

func countPossibleTimes(time int, record int) int {
	// T ms D mm
	// x mm/ms
	// (T-x) * x
	// Tx - x^2 > D
	// -x^2 + Tx - D = 0
	// d = T^2 - 4 * -1 * - D = T^2 - 4D
	// x1 = (-T + sqrt(T^2 - 4D)) / -2 = T/2 - sqrt(T^2 - 4D)/2
	// x2 = (-T - sqrt(T^2 - 4D)) / -2 = T/2 + sqrt(T^2 - 4D)/2
	// T/2 - sqrt(T^2 - 4D)/2 < x < T/2 + sqrt(T^2 - 4D)/2
	f_time := float64(time)
	f_record := float64(record)
	s_delta := math.Sqrt(f_time*f_time - 4.0*f_record)
	f_x1 := (f_time - s_delta) / 2
	f_x2 := (f_time + s_delta) / 2
	var x1, x2 int
	if f_x1 == math.Ceil(f_x1) {
		x1 = int(f_x1) + 1
	} else {
		x1 = int(math.Ceil(f_x1))
	}

	if f_x2 == math.Floor(f_x2) {
		x2 = int(f_x2) - 1
	} else {
		x2 = int(math.Floor(f_x2))
	}

	return x2 - x1 + 1
}

func processLines(path *string) int {
	f, e := os.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	timeValues := getValues(scanner)
	distanceValues := getValues(scanner)
	prod := 1
	for i := range timeValues {
		prod *= countPossibleTimes(timeValues[i], distanceValues[i])
	}

	return prod
}

func main() {
	path := "input.txt"
	output := processLines(&path)
	fmt.Println(output)
}
