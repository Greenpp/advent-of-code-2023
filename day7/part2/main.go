package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	sc "strconv"
	s "strings"
)

type Hand struct {
	Repr string
	Bid  int
}

type ByRepr []Hand

func (a ByRepr) Len() int           { return len(a) }
func (a ByRepr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRepr) Less(i, j int) bool { return a[i].Repr < a[j].Repr }

func buildCardsMap() map[rune]rune {
	cards := "J23456789TQKA"
	cardMap := map[rune]rune{}
	for i, c := range cards {
		cardMap[c] = rune(97 + i)
	}
	return cardMap
}

func mapCards(cards *string, cardsMap map[rune]rune) []rune {
	var mapped []rune
	for _, c := range *cards {
		mapped = append(mapped, cardsMap[c])
	}
	return mapped
}

func mapHand(cards *string) rune {
	cardsCount := map[rune]int{}
	jokers := 0
	for _, c := range *cards {
		if c != 'J' {
			cardsCount[c] += 1
		} else {
			jokers += 1
		}
	}
	counts := make([]int, 0, len(cardsCount))
	for _, c := range cardsCount {
		counts = append(counts, c)
	}
	if len(counts) > 0 {
		sort.Ints(counts)
		counts[len(counts)-1] += jokers
	} else {
		counts = append(counts, jokers)
	}

	if len(counts) == 5 {
		// High card
		return rune(1)
	}
	if len(counts) == 4 {
		// One pair
		return rune(2)
	}
	if len(counts) == 1 {
		// Five of a kind
		return rune(7)
	}
	if len(counts) == 2 {
		if counts[0] == 4 || counts[1] == 4 {
			// Four of a kind
			return rune(6)
		} else {
			// Full house
			return rune(5)
		}
	}
	if len(counts) == 3 {
		if counts[0] == 3 || counts[1] == 3 || counts[2] == 3 {
			// Three of a kind
			return rune(4)
		} else {
			// Two pair
			return rune(3)
		}
	}

	log.Fatalf("Failed to parse hand %s", *cards)
	return 'a'
}

func buildHandRepr(hand *string, cardsMap map[rune]rune) *string {
	cardsRepr := mapCards(hand, cardsMap)
	handRepr := mapHand(hand) + 97

	reprSlice := append([]rune{handRepr}, cardsRepr...)
	repr := string(reprSlice)
	return &repr
}

func processLine(line *string, cardsMap map[rune]rune) Hand {
	parts := s.Split(*line, " ")
	raw_hand := parts[0]
	bid, e := sc.Atoi(parts[1])
	if e != nil {
		log.Fatalf("Failed to convert %s to int", parts[1])
	}
	return Hand{
		Bid:  bid,
		Repr: *buildHandRepr(&raw_hand, cardsMap),
	}

}

func processLines(path *string) int {
	f, e := os.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	cardsMap := buildCardsMap()

	scanner := bufio.NewScanner(f)
	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		hands = append(hands, processLine(&line, cardsMap))
	}

	sort.Sort(ByRepr(hands))
	sum := 0
	for i, h := range hands {
		sum += (h.Bid * (i + 1))
	}
	return sum
}

func main() {
	path := "input.txt"
	output := processLines(&path)
	fmt.Println(output)
}
