package main

import (
	"fmt"
	"math"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

type inputT elvesMapT

type elvesMapT map[positionT]int

type positionT struct {
	row, col int
}

type directionT []positionT

var N = directionT{{-1, -1}, {-1, 0}, {-1, 1}}
var S = directionT{{1, -1}, {1, 0}, {1, 1}}
var E = directionT{{-1, 1}, {0, 1}, {1, 1}}
var W = directionT{{-1, -1}, {0, -1}, {1, -1}}

func isFree(elvesMap elvesMapT, elfPosition positionT, direction directionT) bool {
	for _, delta := range direction {
		position := positionT{elfPosition.row + delta.row, elfPosition.col + delta.col}
		if elvesMap[position] > 0 {
			return false
		}
	}
	return true
}

func getElfProposal(elvesMap elvesMapT, elf positionT, directions []directionT) positionT {
	allIsFree := true
	var firstFreePosition *positionT

	for _, direction := range directions {
		_isFree := isFree(elvesMap, elf, direction)
		allIsFree = allIsFree && _isFree

		if _isFree && firstFreePosition == nil {
			delta := direction[1]
			firstFreePosition = &positionT{elf.row + delta.row, elf.col + delta.col}
		}
	}

	if allIsFree || firstFreePosition == nil {
		return elf
	}

	return *firstFreePosition
}

func getNextElvesMap(elvesMap elvesMapT, directions []directionT) elvesMapT {
	nextElvesMap := elvesMapT{}
	proposal := map[positionT]positionT{}
	blacklist := elvesMapT{}

	for elf := range elvesMap {
		elfProposal := getElfProposal(elvesMap, elf, directions)
		proposal[elf] = elfProposal
		blacklist[elfProposal]++
	}

	for curr, next := range proposal {
		if blacklist[next] > 1 {
			nextElvesMap[curr] = 1
		} else {
			nextElvesMap[next] = 1
		}
	}

	return nextElvesMap
}

func isEqual(map1 elvesMapT, map2 elvesMapT) bool {
	if len(map1) != len(map2) {
		return false
	}

	for k1 := range map1 {
		if _, ok := map2[k1]; !ok {
			return false
		}
	}

	return true
}

func getFinalElvesMap(initialElvesMap elvesMapT, nRounds int) (elvesMapT, int) {
	if nRounds == 0 {
		return initialElvesMap, nRounds
	}

	currMap := elvesMapT{}
	nextMap := initialElvesMap
	directions := []directionT{N, S, W, E}

	n := 0

	for (n <= nRounds || nRounds == -1) && !isEqual(currMap, nextMap) {
		currMap = nextMap
		countFreeTiles(currMap)

		nextMap = getNextElvesMap(currMap, directions)
		directions = append(directions[1:], directions[0])

		n++
	}

	return currMap, n
}

func countFreeTiles(elvesMap elvesMapT) int {
	var count, minRow, maxRow, minCol, maxCol int

	minRow = math.MaxInt
	minCol = math.MaxInt

	for elf := range elvesMap {
		minRow = utils.Min(minRow, elf.row)
		maxRow = utils.Max(maxRow, elf.row)
		minCol = utils.Min(minCol, elf.col)
		maxCol = utils.Max(maxCol, elf.col)
	}

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if _, ok := elvesMap[positionT{row, col}]; !ok {
				count++
			}
		}
	}

	return count
}

func part1(input inputT) {
	finalMap, _ := getFinalElvesMap(elvesMapT(input), 10)
	answer := countFreeTiles(finalMap)
	fmt.Println("part 1:", answer)
}

func part2(input inputT) {
	_, finalRound := getFinalElvesMap(elvesMapT(input), -1)
	answer := finalRound
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
