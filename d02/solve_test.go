package main

import "testing"

func TestGetRoundScore(t *testing.T) {
	tcs := []struct {
		opShape string
		myShape string
		want    int
	}{
		{"A", "Y", 6},
		{"B", "X", 0},
		{"C", "Z", 3},
		{"A", "Z", 0},
		{"C", "X", 6},
	}

	for _, tc := range tcs {
		got := getRoundScore(tc.opShape, tc.myShape)
		want := tc.want

		if got != want {
			t.Errorf("want: %v, got: %v", want, got)
		}
	}
}

func TestGetTotalScore(t *testing.T) {
	rounds := parseFile("example.txt")

	want := 15
	got := getTotalScore(rounds)

	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
