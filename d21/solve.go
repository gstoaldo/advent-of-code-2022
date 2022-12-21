package main

import (
	"fmt"
)

type inputT nodeT

type nodeT struct {
	name     string
	operator string
	value    int
	left     *nodeT
	right    *nodeT
}

func runJob(node nodeT) int {
	switch operator := node.operator; operator {
	case "+":
		return runJob(*node.left) + runJob(*node.right)
	case "-":
		return runJob(*node.left) - runJob(*node.right)
	case "*":
		return runJob(*node.left) * runJob(*node.right)
	case "/":
		return runJob(*node.left) / runJob(*node.right)
	default:
		return node.value
	}
}

func part1(input inputT) {
	answer := runJob(nodeT(input))
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
