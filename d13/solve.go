package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

type inputType [][]string

func findStartIndex(s string, targetLevel []int) int {
	var currLevel []int

	for i, r := range s {
		ch := string(r)
		if ch == "[" {
			currLevel = append(currLevel, 0)
		}

		if ch == "," {
			currLevel[len(currLevel)-1]++
		}

		if ch == "]" {
			currLevel = currLevel[:len(currLevel)-1]
		}

		if reflect.DeepEqual(currLevel, targetLevel) {
			return i
		}
	}

	return -1
}

func findCloseIndex(s string, startIndex int) int {
	count := 0

	for i, r := range s[startIndex+1:] {
		ch := string(r)

		if ch == "[" {
			count++
		}

		if ch == "]" {
			count--
		}

		if ch == "," && count == 0 {
			return startIndex + 1 + i
		}

		if count == -1 {
			return startIndex + 1 + i
		}
	}

	return -1
}

func findElementByLevel(s string, targetLevel []int) string {
	// given a string representing a slice, find the element given the "target level"
	// ex:
	// slice: [1, 2, [3, 4]]
	// targetLevel: [2,0] -> return "4", same as slice[2][0]
	// targetLevel: [2] -> return "[3,4]", same as slice[2]

	endTarget := append([]int{}, targetLevel...)
	endTarget[len(endTarget)-1]++

	startIndex := findStartIndex(s, targetLevel)
	endIndex := findCloseIndex(s, startIndex)

	if startIndex == -1 || endIndex == -1 {
		return ""
	}

	return s[startIndex+1 : endIndex]
}

func pairIsOrdered(pL string, pR string) (shouldContinue bool, ordered bool) {
	shouldContinue = true
	ordered = true
	i := 0

	for shouldContinue {
		left := findElementByLevel(pL, []int{i})
		right := findElementByLevel(pR, []int{i})

		leftValue, errLeft := strconv.Atoi(left)
		rightValue, errRight := strconv.Atoi(right)

		if left == "" && right == "" {
			return true, true
		}

		// left ends first
		if left == "" && right != "" {
			return false, true
		}

		// right ends first
		if left != "" && right == "" {
			return false, false
		}

		// both are number
		if errLeft == nil && errRight == nil {

			if leftValue < rightValue {
				return false, true
			}

			shouldContinue = leftValue == rightValue
			ordered = shouldContinue
		}

		// left is number, right is list
		if errLeft == nil && errRight != nil {
			left = fmt.Sprintf("[%v]", left)
			shouldContinue, ordered = pairIsOrdered(left, right)
		}

		// left is list, right is number
		if errLeft != nil && errRight == nil {
			right = fmt.Sprintf("[%v]", right)
			shouldContinue, ordered = pairIsOrdered(left, right)
		}

		// both are list
		if errLeft != nil && errRight != nil {
			shouldContinue, ordered = pairIsOrdered(left, right)
		}

		i++
	}

	return shouldContinue, ordered
}

//

func sumOrderedPairIndices(pairs inputType) int {
	sum := 0
	for i, pair := range pairs {
		if _, ordered := pairIsOrdered(pair[0], pair[1]); ordered {
			sum += i + 1
		}
	}

	return sum
}

func sortPackages(packages []string) {
	sort.Slice(packages, func(i, j int) bool {
		_, ordered := pairIsOrdered(packages[i], packages[j])
		return ordered
	})
}

func findDecoderKey(input inputType) int {
	packages := []string{}

	for _, pair := range input {
		packages = append(packages, pair...)
	}

	firstDivider := "[[2]]"
	secondDivider := "[[6]]"

	packages = append(packages, []string{firstDivider, secondDivider}...)

	sortPackages(packages)

	firstDividerId := 0
	secondDividerId := 0

	for i, p := range packages {
		if p == firstDivider {
			firstDividerId = i + 1
		}

		if p == secondDivider {
			secondDividerId = i + 1
		}
	}

	return firstDividerId * secondDividerId
}

func part1(input inputType) {
	answer := sumOrderedPairIndices(input)
	fmt.Println("part 1:", answer)
}

func part2(input inputType) {
	answer := findDecoderKey(input)
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
