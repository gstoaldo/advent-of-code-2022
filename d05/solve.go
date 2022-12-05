package main

import "fmt"

type Input [][]int

func runMove(stacks [][]string, move []int) {
	from := move[1]
	to := move[2]

	stackMoveFrom := stacks[from]
	stackMoveTo := stacks[to]

	crate := stackMoveFrom[len(stackMoveFrom)-1]
	stackMoveTo = append(stackMoveTo, crate)

	stacks[from] = stackMoveFrom[:len(stackMoveFrom)-1]
	stacks[to] = stackMoveTo
}

func runMoves(stacks [][]string, moves [][]int) {
	for _, move := range moves {
		for i := 0; i < move[0]; i++ {
			runMove(stacks, move)
		}
	}
}

func getTopCranes(stacks [][]string) string {
	topCranes := ""

	for _, stack := range stacks {
		topCranes += stack[len(stack)-1]
	}

	return topCranes
}

func part1(input inputType) {
	runMoves(input.stacks, input.moves)
	answer := getTopCranes(input.stacks)
	fmt.Println("part 1:", answer)
}

// func part2(input [][]int) {
// answer := countOverlapPairs(input)
// fmt.Println("part 2:", answer)
// }

func main() {
	input := parseFile("input.txt")

	part1(input)
	// part2(input)
}
