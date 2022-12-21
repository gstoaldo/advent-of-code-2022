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

func containsHuman(node nodeT) bool {
	if node.name == "humn" {
		return true
	}

	isOnTheLeft := false
	if node.left != nil {
		isOnTheLeft = containsHuman(*node.left)
	}

	isOnTheRight := false
	if node.right != nil {
		isOnTheRight = containsHuman(*node.right)
	}

	return isOnTheLeft || isOnTheRight
}

func calcHumanValue(node nodeT) int {
	var knownValue int
	var otherValue int
	var nextNode *nodeT
	onTheLeft := false

	if node.name == "humn" {
		return node.value
	}

	if containsHuman(*node.left) {
		knownValue = runJob(*node.right)
		nextNode = node.left
		onTheLeft = true
	} else {
		knownValue = runJob(*node.left)
		nextNode = node.right
	}

	switch node.operator {
	case "+":
		if node.name == "root" {
			otherValue = knownValue
		} else {
			otherValue = node.value - knownValue
		}
	case "-":
		if onTheLeft {
			otherValue = node.value + knownValue
		} else {
			otherValue = knownValue - node.value
		}
	case "*":
		otherValue = node.value / knownValue
	case "/":
		if onTheLeft {
			otherValue = node.value * knownValue
		} else {
			otherValue = knownValue / node.value
		}
	default:
		panic("operator not found")
	}

	nextNode.value = otherValue
	return calcHumanValue(*nextNode)
}

func part1(input inputT) {
	answer := runJob(nodeT(input))
	fmt.Println("part 1:", answer)
}

func part2(input inputT) {
	answer := calcHumanValue(nodeT(input))
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
