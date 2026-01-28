package main

import "testing"

type testTable struct {
	rotSeq   []string
	expected int
}

var testTables = []testTable{
	{[]string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}, 3},
	{[]string{"R50", "L30", "R13", "L44", "L35", "L4", "R80", "R20", "R14", "L82"}, 3},
}

func TestGetNumOfDialAtZero(t *testing.T) {
	for _, test := range testTables {
		output, err := GetNumOfDialAtZero(test.rotSeq)
		if err != nil {
			t.Errorf("%s", err)
		}
		if output != test.expected {
			t.Errorf("got %d, want %d", output, test.expected)
		}
	}
}
