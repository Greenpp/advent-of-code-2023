package main

import (
	"advent/day1/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	const expected = 142
	path := "test_data.txt"
	sum := utils.ProcessFile(&path, lineToValue)

	if sum != expected {
		t.Logf("Sum should be %d (got %d)", expected, sum)
		t.Fail()
	}
}
