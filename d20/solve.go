package main

import (
	"fmt"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

type inputT []int
type numberT struct{ value, originalId int }
type inputMapT map[numberT]int

func calcNewId(value, oldId, length int) int {
	newId := oldId + value%(length-1)

	if newId >= length && value > 0 {
		return newId%length + 1
	}

	if newId <= 0 && value < 0 {
		return newId + length - 1
	}

	return newId
}

func getInputMap(input inputT) inputMapT {
	inputMap := inputMapT{}

	for id, value := range input {
		inputMap[numberT{value, id}] = id
	}

	return inputMap
}

func move(number numberT, inputMap inputMapT) {
	if number.value == 0 {
		return
	}

	oldId := inputMap[number]
	newId := calcNewId(number.value, oldId, len(inputMap))

	minId := utils.Min(oldId, newId)
	maxId := utils.Max(oldId, newId)

	for number, id := range inputMap {
		if id < minId || id > maxId {
			continue
		}

		if id > oldId {
			inputMap[number]--
		}

		if id < oldId {
			inputMap[number]++
		}
	}

	inputMap[number] = newId
}

func mix(input inputT) []int {
	inputMap := getInputMap(input)

	for id, value := range input {
		move(numberT{value, id}, inputMap)
	}

	final := make([]int, len(inputMap))

	for number, id := range inputMap {
		final[id] = number.value
	}

	return final
}

func getNth(values []int, nth int) int {
	zeroId := -1
	for id, val := range values {
		if val == 0 {
			zeroId = id
		}
	}

	targetId := (zeroId + nth) % len(values)

	return values[targetId]
}

func getGroveCoordinates(values []int) int {
	sum := 0
	for _, n := range []int{1000, 2000, 3000} {
		sum += getNth(values, n)
	}

	return sum
}

func part1(input inputT) {
	final := mix(input)
	answer := getGroveCoordinates(final)
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
