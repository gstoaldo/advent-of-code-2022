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
		got := getRoundScore(tc.opShape, mapColumn[tc.myShape])
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

func TestGetShapeToResult(t *testing.T) {
	tcs := []struct {
		opShape string
		result  string
		want    string
	}{
		{"A", "Y", "A"},
		{"B", "X", "A"},
		{"C", "Z", "A"},
		{"A", "X", "C"},
	}

	for _, tc := range tcs {
		got := getShapeToResult(tc.opShape, tc.result)
		want := tc.want

		if got != want {
			t.Errorf("want: %v, got: %v", want, got)
		}
	}
}
