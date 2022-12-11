package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var example1 = parseFile("example1.txt")

func TestRunRoundsReliefPT1(t *testing.T) {
	tcs := []struct {
		input   []monkey
		nrounds int
		want    [][]int
	}{
		{
			example1.monkeys,
			1,
			[][]int{{20, 23, 27, 26}, {2080, 25, 167, 207, 401, 1046}, {}, {}},
		},
		{
			example1.monkeys,
			2,
			[][]int{{695, 10, 71, 135, 350}, {43, 49, 58, 55, 362}, {}, {}},
		},
		{
			example1.monkeys,
			20,
			[][]int{{10, 12, 14, 26, 34}, {245, 93, 53, 199, 115}, {}, {}},
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			cp := make([]monkey, len(tc.input))
			copy(cp, tc.input)

			relief := reliefPT1

			runRounds(cp, tc.nrounds, relief)

			for i := range cp {
				assert.Equal(t, tc.want[i], cp[i].items)
			}
		})
	}
}

func TestRunRoundsReliefPT2(t *testing.T) {
	tcs := []struct {
		input   inputType
		nrounds int
		want    []int
	}{
		{
			example1,
			1,
			[]int{2, 4, 3, 6},
		},
		{
			example1,
			20,
			[]int{99, 97, 8, 103},
		},
		{
			example1,
			1000,
			[]int{5204, 4792, 199, 5192},
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			cp := make([]monkey, len(tc.input.monkeys))
			copy(cp, tc.input.monkeys)

			relief := getReliefPT2(tc.input.mmc)

			runRounds(cp, tc.nrounds, relief)

			count := []int{}

			for _, m := range cp {
				count = append(count, m.inspectedItemsCount)
			}

			assert.Equal(t, tc.want, count)
		})
	}
}
