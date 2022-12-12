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

func replace(s string) string {
	if s == "S" {
		return "a"
	}

	if s == "E" {
		return "z"
	}

	return s
}

func deltaHeight(grid [][]string, pStart position, pEnd position) int {
	start := replace(grid[pStart.i][pStart.j])
	end := replace(grid[pEnd.i][pEnd.j])

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

		print(grid, visited, end)

		stepQueue = nextStepQueue
		nextStepQueue = []position{}
		step++
	}

	return -1
}

func print(grid [][]string, visited map[position]int, end position) {
	for i, row := range grid {
		for j, s := range row {
			v := s

			if _, ok := visited[position{i, j}]; ok {
				v = "[*]"
			} else {
				v = fmt.Sprintf(" %v ", v)
			}

			if i == end.i && j == end.j {
				v = fmt.Sprintf("{%v}", v)
			}

			fmt.Printf("%v", v)
		}

		fmt.Printf("\n")
	}
}

func part1(input inputType) {
	answer := bfs(input.grid, input.start, input.end)
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
