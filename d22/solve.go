package main

import (
	"fmt"
	"strconv"
)

type inputT struct {
	board    boardT
	commands []string
	start    pointT
}

type pointT struct {
	x, y int
}

type vectorT struct {
	x, y int
}

type boardT map[pointT]bool

var DIRECTIONS = [4]vectorT{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func rotate(facing int, command string) int {
	length := len(DIRECTIONS)

	if command == "R" {
		return (facing + 1) % length
	}

	return (facing - 1 + length) % length
}

func wrap(board boardT, currPosition pointT, v vectorT) pointT {
	nextExist, currIsOpen := true, true
	nextIsOpen := false

	curr := pointT{}
	next := currPosition

	for nextExist {
		curr = next
		currIsOpen = nextIsOpen
		next = pointT{curr.x - v.x, curr.y - v.y}

		nextIsOpen, nextExist = board[next]
	}

	if currIsOpen {
		return curr
	}

	return currPosition
}

func move(board boardT, currPosition pointT, v vectorT, nTiles int) pointT {
	curr := pointT{}
	next := currPosition

	for n := 0; n < nTiles; n++ {
		curr = next
		next = pointT{curr.x + v.x, curr.y + v.y}

		isOpen, exist := board[next]

		if !isOpen {
			next = curr
		}

		if !exist {
			next = wrap(board, curr, v)
		}

		if next == curr {
			break
		}
	}

	return next
}

func simulate(board boardT, commands []string, start pointT, facing int) (pointT, int) {
	currPosition := start

	for _, command := range commands {
		nTiles, err := strconv.Atoi(command)

		if err == nil {
			currPosition = move(board, currPosition, DIRECTIONS[facing], nTiles)
		} else {
			facing = rotate(facing, command)
		}
	}

	return currPosition, facing
}

func calcPassword(row, column, facing int) int {
	return 1000*row + 4*column + facing
}

func part1(input inputT) {
	finalPosition, facing := simulate(input.board, input.commands, input.start, 0)
	answer := calcPassword(finalPosition.y, finalPosition.x, facing)
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
