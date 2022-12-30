package main

import (
	"fmt"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

type inputT []blueprintT

type blueprintT [][]int

type stateT struct {
	resources []int // ore, clay, obs, geo
	robots    []int
	minute    int
}

func canBuildNow(state stateT, cost []int) bool {
	for i, amount := range state.resources {
		if amount < cost[i] {
			return false
		}
	}
	return true
}

func canBuildInTheFuture(state stateT, cost []int) bool {
	for i, resourceCost := range cost {
		if state.robots[i] == 0 && resourceCost != 0 {
			return false
		}
	}
	return true
}

func doNothing(state stateT) stateT {
	nextState := stateT{
		resources: append([]int{}, state.resources...),
		robots:    append([]int{}, state.robots...),
		minute:    state.minute + 1,
	}

	for i := range state.resources {
		nextState.resources[i] = state.resources[i] + state.robots[i]
	}

	return nextState
}

func buildNow(state stateT, cost []int, robotId int) stateT {
	nextState := stateT{
		resources: append([]int{}, state.resources...),
		robots:    append([]int{}, state.robots...),
		minute:    state.minute + 1,
	}

	for i := range state.resources {
		nextState.resources[i] = state.resources[i] + state.robots[i]
	}

	if !canBuildNow(state, cost) {
		panic("not enough resources to build robot")
	}

	nextState.robots[robotId]++

	for i := range cost {
		nextState.resources[i] -= cost[i]
	}

	return nextState
}

func simulateBlueprint(bp blueprintT, maxTime int) []stateT {
	initialState := stateT{
		resources: []int{0, 0, 0, 0},
		robots:    []int{1, 0, 0, 0},
		minute:    0,
	}

	visited := map[string][][]int{}
	queue := []stateT{initialState}

	addToQueue := func(state stateT) {
		// for a given minute, states that have the same amount of robots but
		// fewer of all resources, should not be added to the queue.

		key := fmt.Sprintf("%v, %v", state.minute, state.robots)
		allResources, ok := visited[key]

		if !ok {
			queue = append(queue, state)
			visited[key] = append(visited[key], state.resources)
			return
		}

		for _, resource := range allResources {
			fewer := true
			for i, r := range resource {
				fewer = fewer && state.resources[i] <= r
			}

			if fewer {
				return
			}
		}

		queue = append(queue, state)
		visited[key] = append(visited[key], state.resources)
	}

	maxConsuptionRate := make([]int, 4)

	for _, robotCost := range bp {
		for i, value := range robotCost {
			maxConsuptionRate[i] = utils.Max(maxConsuptionRate[i], value)
		}
	}

	for queue[0].minute < maxTime {
		state := queue[0]
		queue = queue[1:]
		// fmt.Printf("len: %v, min: %v,robots: %v, resources: %v\n", len(queue), state.minute, state.robots, state.resources)

		shouldWait := false

		for robotId := range bp {
			cost := bp[robotId]

			if state.robots[robotId] == maxConsuptionRate[robotId] && robotId < 3 {
				// we can build only one robot per minute, so it doesnt make sense
				// to have more robots than the consuption rate.
				// except for the geode robot, that we want to build as much as we can.
				continue
			}

			if canBuildNow(state, cost) {
				nextState := buildNow(state, cost, robotId)
				addToQueue(nextState)
				continue
			}

			shouldWait = shouldWait || canBuildInTheFuture(state, cost)
		}

		if shouldWait {
			doNothingState := doNothing(state)
			addToQueue(doNothingState)
		}
	}

	return queue
}

func getMaxGeode(states []stateT) int {
	max := 0
	for _, state := range states {
		max = utils.Max(max, state.resources[3])
	}

	return max
}

func getQualityLevelSum(blueprints inputT) int {
	sum := 0
	for i, bp := range blueprints {
		states := simulateBlueprint(bp, 24)
		maxGeode := getMaxGeode(states)
		sum += (i + 1) * maxGeode
	}

	return sum
}

func getBlueprintProduct(blueprints inputT, n int) int {
	product := 1
	for _, bp := range blueprints[:n] {
		states := simulateBlueprint(bp, 32)
		maxGeode := getMaxGeode(states)
		product *= maxGeode
	}

	return product
}

func part1(input inputT) {
	answer := getQualityLevelSum(input)
	fmt.Println("part 1:", answer)
}

func part2(input inputT) {
	answer := getBlueprintProduct(input, 3)
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
