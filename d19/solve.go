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

func canAssembleNow(state stateT, cost []int) bool {
	for i, amount := range state.resources {
		if amount < cost[i] {
			return false
		}
	}
	return true
}

func canAssembleInTheFuture(state stateT, cost []int) bool {
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

func assembleNow(state stateT, cost []int, robotId int) stateT {
	nextState := stateT{
		resources: append([]int{}, state.resources...),
		robots:    append([]int{}, state.robots...),
		minute:    state.minute + 1,
	}

	for i := range state.resources {
		nextState.resources[i] = state.resources[i] + state.robots[i]
	}

	if !canAssembleNow(state, cost) {
		panic("not enough resources to assemble robot")
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

	visited := map[string]bool{}
	queue := []stateT{initialState}

	addToQueue := func(state stateT) {
		key := fmt.Sprintf("%v, %v, %v", state.minute, state.robots, state.resources)

		if !visited[key] {
			queue = append(queue, state)
			visited[key] = true
		}
	}

	resourceMax := make([]int, 4)

	for _, robotCost := range bp {
		for i, value := range robotCost {
			resourceMax[i] = utils.Max(resourceMax[i], value)
		}
	}

	for queue[0].minute < maxTime {
		state := queue[0]
		queue = queue[1:]
		// fmt.Printf("len: %v, min: %v,robots: %v, resources: %v\n", len(queue), state.minute, state.robots, state.resources)

		shouldWait := false

		for robotId := range bp {
			cost := bp[robotId]

			if state.robots[robotId] == resourceMax[robotId] && robotId < 3 {
				continue
			}

			if canAssembleNow(state, cost) {
				nextState := assembleNow(state, cost, robotId)
				addToQueue(nextState)

				continue
			}

			shouldWait = shouldWait || canAssembleInTheFuture(state, cost)
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
		fmt.Println("bp:", i)
		states := simulateBlueprint(bp, 24)
		maxGeode := getMaxGeode(states)
		sum += (i + 1) * maxGeode
	}

	return sum
}

func part1(input inputT) {
	answer := getQualityLevelSum(input)
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
