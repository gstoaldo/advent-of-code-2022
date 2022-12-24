package main

import (
	"fmt"
	"math"
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

type edgeT struct{ face, facing int }

type wrapFunc func(boardT, pointT, int) (pointT, int)

var exampleFaceMap = map[edgeT]struct{ x0, y0, angle, face1, facing1 int }{
	{1, 0}: {12, 1, -90, 6, 0},
	{1, 2}: {9, 1, -90, 3, 3},
	{1, 3}: {12, 1, 180, 2, 3},
	{2, 1}: {4, 8, 180, 5, 1},
	{2, 2}: {1, 8, 90, 6, 1},
	{2, 3}: {1, 5, 0, 1, 3},
	{3, 1}: {5, 8, 0, 5, 2},
	{3, 3}: {5, 5, 0, 1, 2},
	{4, 0}: {12, 8, 90, 6, 3},
	{5, 1}: {9, 12, 0, 2, 1},
	{5, 2}: {9, 12, 90, 3, 1},
	{6, 0}: {16, 12, 90, 1, 0},
	{6, 1}: {13, 12, 0, 2, 2},
	{6, 3}: {13, 9, 0, 4, 0},
}

var inputFaceMap = map[edgeT]struct{ x0, y0, angle, face1, facing1 int }{
	{1, 2}: {51, 1, -90, 4, 2},
	{1, 3}: {51, 1, 0, 6, 2},
	{2, 0}: {150, 50, 90, 5, 0},
	{2, 1}: {101, 50, 0, 3, 0},
	{2, 3}: {101, 1, 0, 6, 1},
	{3, 0}: {100, 51, -90, 2, 1},
	{3, 2}: {51, 51, -90, 4, 3},
	{4, 2}: {1, 150, 90, 1, 2},
	{4, 3}: {1, 101, 0, 3, 2},
	{5, 0}: {100, 101, -90, 2, 0},
	{5, 1}: {51, 150, 0, 6, 0},
	{6, 0}: {50, 151, -90, 5, 1},
	{6, 1}: {1, 200, 0, 2, 3},
	{6, 2}: {1, 151, -90, 1, 3},
}

var CUBE_SIDE = 4
var CUBE_FACE_MAP = exampleFaceMap

//

func rotate(facing int, command string) int {
	length := len(DIRECTIONS)

	if command == "R" {
		return (facing + 1) % length
	}

	return (facing - 1 + length) % length
}

func wrap2D(board boardT, initialPosition pointT, facing int) (pointT, int) {
	v := DIRECTIONS[facing]
	nextExist := true

	curr := pointT{}
	next := initialPosition

	for nextExist {
		curr = next
		next = pointT{curr.x - v.x, curr.y - v.y}

		_, nextExist = board[next]
	}

	return curr, facing
}

func move(board boardT, currPosition pointT, facing int, nTiles int, wrap wrapFunc) (pointT, int) {
	curr := pointT{}
	next := currPosition
	nextFacing := facing

	for n := 0; n < nTiles; n++ {
		curr = next
		facing = nextFacing
		v := DIRECTIONS[facing]
		next = pointT{curr.x + v.x, curr.y + v.y}

		isOpen, exist := board[next]

		if !isOpen {
			next = curr
		}

		if !exist {
			next, nextFacing = wrap(board, curr, facing)
			if !board[next] {
				next = curr
				nextFacing = facing
			}
		}

		if next == curr {
			break
		}
	}

	return next, facing
}

func simulate(board boardT, commands []string, start pointT, facing int, wrap wrapFunc) (pointT, int) {
	currPosition := start

	for _, command := range commands {
		nTiles, err := strconv.Atoi(command)

		if err == nil {
			currPosition, facing = move(board, currPosition, facing, nTiles, wrap)
		} else {
			facing = rotate(facing, command)
		}
	}

	return currPosition, facing
}

func calcPassword(row, column, facing int) int {
	return 1000*row + 4*column + facing
}

//

func convertCoords(p0 pointT, face, facing int) (pointT, int) {
	face0 := CUBE_FACE_MAP[edgeT{face, facing}]
	face1 := CUBE_FACE_MAP[edgeT{face0.face1, face0.facing1}]

	relativePoint := pointT{p0.x - face0.x0, p0.y - face0.y0}
	angle := face0.angle - face1.angle

	angleRad := float64(angle) * math.Pi / 180.0
	cos := int(math.Cos(angleRad))
	sin := int(math.Sin(angleRad))

	relativeRotated := pointT{cos*relativePoint.x - sin*relativePoint.y, sin*relativePoint.x + cos*relativePoint.y}
	absolute := pointT{face1.x0 + relativeRotated.x, face1.y0 + relativeRotated.y}

	return absolute, reverse(face0.facing1)
}

func currentFace(board boardT, currPoint pointT, side int) int {
	face := 1
	for y := currPoint.y % side; y < 6*side; y += side {
		for x := currPoint.x % side; x < 6*side; x += side {
			point := pointT{x, y}
			_, exist := board[point]

			if currPoint == point && exist {
				return face
			}

			if exist {
				face++
			}
		}
	}

	return face
}

func reverse(facing int) int {
	return (facing + 2) % 4
}

func wrapCube(board boardT, curr pointT, facing int) (pointT, int) {
	face := currentFace(board, curr, CUBE_SIDE)
	next, nextFacing := convertCoords(curr, face, facing)

	return next, nextFacing
}

//

func part1(input inputT) {
	finalPosition, facing := simulate(input.board, input.commands, input.start, 0, wrap2D)
	answer := calcPassword(finalPosition.y, finalPosition.x, facing)
	fmt.Println("part 1:", answer)
}

func part2(input inputT) {
	CUBE_SIDE = 4
	// CUBE_SIDE = 50

	CUBE_FACE_MAP = exampleFaceMap
	// CUBE_FACE_MAP = inputFaceMap

	finalPosition, facing := simulate(input.board, input.commands, input.start, 0, wrapCube)
	answer := calcPassword(finalPosition.y, finalPosition.x, facing)
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("example.txt")

	part1(input)
	part2(input) // working for example but failing for input :/
}
