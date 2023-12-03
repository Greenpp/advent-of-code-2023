package main

import (
	"advent/day1/utils"
	"testing"
)

func TestPart2(t *testing.T) {
	const expected = 281
	path := "test_data.txt"
	output := utils.ProcessFile(&path, lineToValue)

	if output != expected {
		t.Logf("Sum should be %d (got %d)", expected, output)
		t.Fail()
	}
}
