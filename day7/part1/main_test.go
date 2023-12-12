package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	const expected = 6440
	path := "test_data.txt"
	output := processLines(&path)

	if output != expected {
		t.Errorf("Expected to get %d but got %d instead", expected, output)
	}

}
