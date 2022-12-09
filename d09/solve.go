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

func isTouching(head position, tail position) bool {
	return math.Abs(float64(head.i-tail.i)) <= 1 && math.Abs(float64(head.j-tail.j)) <= 1
}

func move(startPosition position, di int, dj int) position {
	return position{startPosition.i + di, startPosition.j + dj}
}

func calcTailMove(head position, tail position) position {
	delta := position{head.i - tail.i, head.j - tail.j}

	diNormalized := 0
	if delta.i != 0 {
		diNormalized = delta.i / int(math.Abs(float64(delta.i)))
	}

	djNormalized := 0
	if delta.j != 0 {
		djNormalized = delta.j / int(math.Abs(float64(delta.j)))
	}

	return position{diNormalized, djNormalized}
}

func moveHeadAndTail(headStart position, tailStart position, di int, dj int) (position, position) {
	headEnd := move(headStart, di, dj)

	if isTouching(headEnd, tailStart) {
		return headEnd, tailStart
	}

	tailMove := calcTailMove(headEnd, tailStart)
	tailEnd := move(tailStart, tailMove.i, tailMove.j)

	return headEnd, tailEnd
}

func getStepMove(step stepType) position {
	moveMap := map[string]position{
		"R": {0, 1},
		"L": {0, -1},
		"U": {1, 0},
		"D": {-1, 0},
	}

	return moveMap[step.direction]
}

func runSteps(steps inputType) int {
	head := position{0, 0}
	tail := position{0, 0}

	tailPathSet := map[position]bool{
		tail: true,
	}

	for _, step := range steps {
		movement := getStepMove(step)
		for i := 0; i < step.value; i++ {
			head, tail = moveHeadAndTail(head, tail, movement.i, movement.j)
			tailPathSet[tail] = true
		}
	}

	return len(tailPathSet)
}

func part1(input inputType) {
	answer := runSteps(input)
	fmt.Println("part 1:", answer)
}

func part2(input inputType) {
	answer := ""
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
