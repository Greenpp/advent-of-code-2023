package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	sc "strconv"
	s "strings"
)

func processLine(line *string) int {
	data := s.Split(*line, ":")[1]
	games := s.Split(data, ";")

	rMax, gMax, bMax := 0, 0, 0
	for _, game := range games {
		tokens := s.Split(game, ",")
		for _, token := range tokens {
			splitToken := s.Split(s.TrimSpace(token), " ")
			color := splitToken[1]
			value, e := sc.Atoi(splitToken[0])
			if e != nil {
				log.Fatalf("Failed to convert %s to integer", splitToken[0])
			}
			switch color {
			case "red":
				rMax = max(rMax, value)
			case "green":
				gMax = max(gMax, value)
			case "blue":
				bMax = max(bMax, value)
			default:
				log.Fatalf("Unknown color %s", color)

			}
		}
	}
	return rMax * gMax * bMax
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
		linePower := processLine(&line)
		sum += linePower
	}

	return sum
}

func main() {
	path := "input.txt"
	output := processLines(&path)
	fmt.Println(output)
}
