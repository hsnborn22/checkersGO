package main

import (
	"reflect"
	"testing"
)

func TestTableCalculateMoves(t *testing.T) {
	game1 := initGame(initBoard())
	var tests = []struct {
		input    int
		input2   int
		expected [][2]int
	}{
		{2, 1, [][2]int{}},
		{0, 0, [][2]int{}},
		{0, 1, [][2]int{}},
		{0, 2, [][2]int{}},
		{0, 3, [][2]int{}},
		{0, 4, [][2]int{}},
		{0, 5, [][2]int{}},
		{0, 6, [][2]int{}},
		{0, 7, [][2]int{}},
		{1, 0, [][2]int{}},
		{1, 0, [][2]int{}},
		{1, 1, [][2]int{}},
		{1, 3, [][2]int{}},
		{1, 5, [][2]int{}},
		{1, 7, [][2]int{}},
		{1, 2, [][2]int{}},
		{1, 4, [][2]int{}},
		{1, 6, [][2]int{}},
		{2, 0, [][2]int{}},
		{2, 1, [][2]int{{3, 0}, {3, 2}}},
		{2, 2, [][2]int{}},
		{2, 3, [][2]int{{2, 2}, {2, 4}}},
		{2, 4, [][2]int{}},
		{2, 5, [][2]int{{2, 4}, {2, 6}}},
		{2, 6, [][2]int{}},
		{2, 7, [][2]int{{2, 6}}},
		{3, 0, [][2]int{}},
		{3, 1, [][2]int{}},
		{3, 2, [][2]int{}},
		{3, 3, [][2]int{}},
		{3, 4, [][2]int{}},
		{3, 5, [][2]int{}},
		{3, 6, [][2]int{}},
		{3, 7, [][2]int{}},
		{4, 0, [][2]int{}},
		{4, 1, [][2]int{}},
		{4, 2, [][2]int{}},
		{4, 3, [][2]int{}},
		{4, 4, [][2]int{}},
		{4, 5, [][2]int{}},
		{4, 6, [][2]int{}},
		{4, 7, [][2]int{}},
		{5, 0, [][2]int{{4, 1}}},
		{5, 1, [][2]int{}},
		{5, 3, [][2]int{}},
		{5, 5, [][2]int{}},
		{5, 7, [][2]int{}},
		{5, 2, [][2]int{{4, 1}, {4, 3}}},
		{5, 4, [][2]int{{4, 3}, {4, 5}}},
		{5, 6, [][2]int{{4, 5}, {4, 7}}},
		{6, 0, [][2]int{}},
		{6, 1, [][2]int{}},
		{6, 2, [][2]int{}},
		{6, 3, [][2]int{}},
		{6, 4, [][2]int{}},
		{6, 5, [][2]int{}},
		{6, 6, [][2]int{}},
		{6, 7, [][2]int{}},
		{7, 0, [][2]int{}},
		{7, 1, [][2]int{}},
		{7, 2, [][2]int{}},
		{7, 3, [][2]int{}},
		{7, 4, [][2]int{}},
		{7, 5, [][2]int{}},
		{7, 6, [][2]int{}},
		{7, 7, [][2]int{}},
	}

	for _, test := range tests {
		if reflect.DeepEqual(calculatePossibleMoves(&game1, test.input, test.input2), test.expected) {
			t.Error("Test Failed: {} and {} inputted, {} expected, received: {}", test.input, test.input2, test.expected, calculatePossibleMoves(&game1, test.input, test.input2))
		}
	}
}
