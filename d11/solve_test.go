package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var example1 = parseFile("example1.txt")

func TestRunRounds(t *testing.T) {
	tcs := []struct {
		monkeys inputType
		nrounds int
		want    [][]int
	}{
		{
			example1,
			1,
			[][]int{{20, 23, 27, 26}, {2080, 25, 167, 207, 401, 1046}, {}, {}},
		},
		{
			example1,
			2,
			[][]int{{695, 10, 71, 135, 350}, {43, 49, 58, 55, 362}, {}, {}},
		},
		{
			example1,
			20,
			[][]int{{10, 12, 14, 26, 34}, {245, 93, 53, 199, 115}, {}, {}},
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			monkeysCopy := make(inputType, len(tc.monkeys))
			copy(monkeysCopy, tc.monkeys)

			runRounds(monkeysCopy, tc.nrounds)

			for i := range tc.monkeys {
				assert.Equal(t, tc.want[i], monkeysCopy[i].items)
			}
		})
	}
}
