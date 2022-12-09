package main

import (
	"fmt"
	"math"
)

type stepType struct {
	direction string
	value     int
}

type inputType []stepType

type position struct {
	i int
	j int
}

type vector position

func isTouching(head position, tail position) bool {
	return math.Abs(float64(head.i-tail.i)) <= 1 && math.Abs(float64(head.j-tail.j)) <= 1
}

func moveByVector(startPosition position, v vector) position {
	return position{startPosition.i + v.i, startPosition.j + v.j}
}

func calcTailVector(head position, tail position) vector {
	delta := position{head.i - tail.i, head.j - tail.j}

	diNormalized := 0
	if delta.i != 0 {
		diNormalized = delta.i / int(math.Abs(float64(delta.i)))
	}

	djNormalized := 0
	if delta.j != 0 {
		djNormalized = delta.j / int(math.Abs(float64(delta.j)))
	}

	return vector{diNormalized, djNormalized}
}

func moveTail(head position, tailStart position) position {
	if isTouching(head, tailStart) {
		return tailStart
	}

	vector := calcTailVector(head, tailStart)
	tailEnd := moveByVector(tailStart, vector)

	return tailEnd
}

func getStepVector(step stepType) vector {
	moveMap := map[string]vector{
		"R": {0, 1},
		"L": {0, -1},
		"U": {1, 0},
		"D": {-1, 0},
	}

	return moveMap[step.direction]
}

func simulate(steps inputType, nknots int) int {
	knots := []position{}

	for i := 0; i < nknots; i++ {
		knots = append(knots, position{0, 0})
	}

	tail := knots[len(knots)-1]

	tailPathSet := map[position]bool{
		tail: true,
	}

	for _, step := range steps {
		vector := getStepVector(step)
		for i := 0; i < step.value; i++ {
			knots[0] = moveByVector(knots[0], vector)

			for i := 1; i < len(knots); i++ {
				knots[i] = moveTail(knots[i-1], knots[i])

				if i == len(knots)-1 {
					tailPathSet[knots[i]] = true
				}
			}
		}
	}

	return len(tailPathSet)
}

func part1(input inputType) {
	answer := simulate(input, 2)
	fmt.Println("part 1:", answer)
}

func part2(input inputType) {
	answer := simulate(input, 10)
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
