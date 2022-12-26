package main

import (
	"fmt"
)

type inputT struct {
	start     positionT
	end       positionT
	blizzards []blizzardT
	maxI      int
	maxJ      int
}

type positionT struct {
	i, j int
}

type velocityT struct {
	i, j int
}

type blizzardT struct {
	p0 positionT
	v0 velocityT
}

func positionEq(p0, v, t, max int) int {
	p := p0 + v*t

	if p >= max {
		return (p-1)%(max-1) + 1
	}

	if p <= 0 {
		return p%(max-1) + max - 1
	}

	return p
}

func blizzardEq(blizzard blizzardT, t, maxI, maxJ int) positionT {
	return positionT{
		i: positionEq(blizzard.p0.i, blizzard.v0.i, t, maxI),
		j: positionEq(blizzard.p0.j, blizzard.v0.j, t, maxJ),
	}
}

func getState(blizzards []blizzardT, t, maxI, maxJ int) map[positionT]int {
	state := map[positionT]int{}

	for _, blizzard := range blizzards {
		p := blizzardEq(blizzard, t, maxI, maxJ)
		state[p]++
	}

	return state
}

func inBounds(p, start, end positionT, maxI, maxJ int) bool {
	_inBounds := p.i > 0 && p.i < maxI && p.j > 0 && p.j < maxJ
	return _inBounds || p == start || p == end
}

func getMoveOptions(state map[positionT]int, currPosition, start, end positionT, maxI, maxJ int) []positionT {
	moveOptions := []positionT{}

	for _, delta := range [][]int{{0, 0}, {1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		p := positionT{currPosition.i + delta[0], currPosition.j + delta[1]}

		if state[p] == 0 && inBounds(p, start, end, maxI, maxJ) {
			moveOptions = append(moveOptions, p)
		}
	}

	return moveOptions
}

func bfs(blizzards []blizzardT, start, end positionT, t0, maxI, maxJ int) int {
	type positionTime struct {
		head positionT
		t    int
	}

	queue := []positionTime{{start, t0}}
	visited := map[positionTime]bool{}

	for len(queue) > 0 {
		head, t := queue[0].head, queue[0].t
		queue = queue[1:]

		if head == end {
			return t
		}

		state := getState(blizzards, t+1, maxI, maxJ)
		moveOptions := getMoveOptions(state, head, start, end, maxI, maxJ)

		for _, moveOption := range moveOptions {
			pt := positionTime{moveOption, t + 1}

			if !visited[pt] {
				visited[pt] = true
				queue = append(queue, positionTime{moveOption, t + 1})
			}
		}
	}

	return -1
}

func threeWayTrip(input inputT) int {
	t1 := bfs(input.blizzards, input.start, input.end, 0, input.maxI, input.maxJ)
	t2 := bfs(input.blizzards, input.end, input.start, t1, input.maxI, input.maxJ)
	t3 := bfs(input.blizzards, input.start, input.end, t2, input.maxI, input.maxJ)

	return t3
}

func part1(input inputT) {
	answer := bfs(input.blizzards, input.start, input.end, 0, input.maxI, input.maxJ)
	fmt.Println("part 1:", answer)
}

func part2(input inputT) {
	answer := threeWayTrip(input)
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
