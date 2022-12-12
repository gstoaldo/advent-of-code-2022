package main

import (
	"fmt"
)

type inputType struct {
	grid  [][]string
	start position
	end   position
}

type position struct {
	i int
	j int
}

func inGrid(grid [][]string, p position) bool {
	return p.i >= 0 && p.i < len(grid) && p.j >= 0 && p.j < len(grid[0])
}

func deltaHeight(grid [][]string, pStart position, pEnd position) int {
	start := grid[pStart.i][pStart.j]
	end := grid[pEnd.i][pEnd.j]

	return int(float64(int(end[0]) - int(start[0])))
}

func findNeighbors(grid [][]string, p position) []position {
	neighbours := []position{}

	for _, delta := range []position{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		n := position{p.i + delta.i, p.j + delta.j}

		if inGrid(grid, n) && deltaHeight(grid, p, n) <= 1 {
			neighbours = append(neighbours, n)
		}
	}

	return neighbours
}

func bfs(grid [][]string, start position, end position) int {
	queue := []struct {
		p     position
		steps int
	}{{start, 0}}

	visited := map[position]bool{}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if _, ok := visited[head.p]; ok {
			continue
		}

		if head.p == end {
			return head.steps
		}

		visited[head.p] = true

		neighbours := findNeighbors(grid, head.p)

		for _, neighbour := range neighbours {
			queue = append(queue, struct {
				p     position
				steps int
			}{neighbour, head.steps + 1})
		}
	}

	return -1
}

func findStartCandidates(grid [][]string) []position {
	candidates := []position{}
	for i, row := range grid {
		for j, r := range row {
			if string(r) == "a" {
				candidates = append(candidates, position{i, j})
			}
		}
	}

	return candidates
}

func findBestStart(grid [][]string, end position) int {
	startCandidates := findStartCandidates(grid)

	minSteps := 999

	for _, start := range startCandidates {
		nsteps := bfs(grid, start, end)

		if nsteps < minSteps && nsteps >= 0 {
			minSteps = nsteps
		}
	}

	return minSteps
}

func part1(input inputType) {
	answer := bfs(input.grid, input.start, input.end)
	fmt.Println("part 1:", answer)
}

func part2(input inputType) {
	answer := findBestStart(input.grid, input.end)
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
