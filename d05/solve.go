package main

import "fmt"

type Input [][]int

func runMove(originalStacks [][]string, move []int, ncrates int) [][]string {
	stacks := [][]string{}

	for _, stack := range originalStacks {
		stackcp := make([]string, len(stack))
		copy(stackcp, stack)
		stacks = append(stacks, stackcp)
	}

	from := move[1]
	to := move[2]

	stackMoveFrom := stacks[from]
	stackMoveTo := stacks[to]

	crates := stackMoveFrom[len(stackMoveFrom)-ncrates:]
	stackMoveTo = append(stackMoveTo, crates...)

	stacks[from] = stackMoveFrom[:len(stackMoveFrom)-ncrates]
	stacks[to] = stackMoveTo

	return stacks
}

func runMovesCrateMover9000(originalStacks [][]string, moves [][]int) [][]string {
	stacks := originalStacks
	for _, move := range moves {
		for i := 0; i < move[0]; i++ {
			stacks = runMove(stacks, move, 1)
		}
	}

	return stacks
}

func runMovesCrateMover9001(originalStacks [][]string, moves [][]int) [][]string {
	stacks := originalStacks
	for _, move := range moves {
		stacks = runMove(stacks, move, move[0])
	}

	return stacks
}

func getTopCranes(stacks [][]string) string {
	topCranes := ""

	for _, stack := range stacks {
		topCranes += stack[len(stack)-1]
	}

	return topCranes
}

func part1(input inputType) {
	stacks := runMovesCrateMover9000(input.stacks, input.moves)
	answer := getTopCranes(stacks)
	fmt.Println("part 1:", answer)
}

func part2(input inputType) {
	stacks := runMovesCrateMover9001(input.stacks, input.moves)
	answer := getTopCranes(stacks)
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
