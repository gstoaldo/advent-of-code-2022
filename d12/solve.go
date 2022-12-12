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
	step := 0
	stepQueue := []position{start}
	nextStepQueue := []position{}
	visited := map[position]int{}
	visited[start] = step

	for n := 0; n < 1000; n++ {
		for len(stepQueue) > 0 {
			head := stepQueue[0]
			stepQueue = stepQueue[1:]

			if head == end {
				return step
			}

			neighbours := findNeighbors(grid, head)

			for _, neighbour := range neighbours {
				dist, ok := visited[neighbour]

				if !ok || (step+1) < dist {
					visited[neighbour] = step + 1
					nextStepQueue = append(nextStepQueue, neighbour)
				}
			}
		}

		stepQueue = nextStepQueue
		nextStepQueue = []position{}
		step++
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
