package main

import (
	"fmt"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

type inputT map[string]nodeT

type nodeT struct {
	flow  int
	nodes []string
}

func getAllPaths(graph inputT) [][]string {
	paths := [][]string{}

	paths = append(paths, []string{"AA"})

	for len(paths[0]) < 29 { // should open at least one valve
		path := paths[0]
		paths = paths[1:]

		head := path[len(path)-1]

		neighbours := graph[head].nodes

		for _, neighbour := range neighbours {
			if len(path) < 2 {
				paths = append(paths, append(path, neighbour))
			} else {
				if neighbour != path[len(path)-2] || len(neighbours) == 1 {
					pathCopy := append([]string{}, path...)
					paths = append(paths, append(pathCopy, neighbour))
				}
			}
		}
	}

	return paths
}

func shouldOpen(idA, idB, flowA, flowB int) bool {
	return (idB-idA+1)*flowA > flowB-1
}

func getPathOptimalValves(graph inputT, path []string) []int {
	openValvesIds := []int{}
	openValvesSet := map[string]bool{}

	for i := 0; i < len(path); i++ {
		curr := path[i]

		open := true

		if graph[curr].flow == 0 || openValvesSet[curr] {
			continue
		}

		for j := i + 1; j < len(path); j++ {
			next := path[j]

			flowA, flowB := graph[curr].flow, graph[next].flow

			open = open && (shouldOpen(i, j, flowA, flowB) || openValvesSet[next])
		}

		if open {
			openValvesIds = append(openValvesIds, i)
			openValvesSet[curr] = true
		}
	}

	return openValvesIds
}

func pressure(graph inputT, path []string, openValvesIds []int) int {
	t := 30

	pressure := 0

	for n, id := range openValvesIds {
		valve := path[id]
		flow := graph[valve].flow
		pressure += (t - (id + 1 + n)) * flow
	}

	return pressure
}

func getMaxPressure(graph inputT) int {
	max := 0

	allPaths := getAllPaths(graph)

	for _, path := range allPaths {
		openIds := getPathOptimalValves(graph, path)
		max = utils.Max(max, pressure(graph, path, openIds))
	}

	return max
}

func part1(input inputT) {
	answer := getMaxPressure(input)
	fmt.Println("part 1:", answer)
}

func part2(input inputT) {
	answer := ""
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
