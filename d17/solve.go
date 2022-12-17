package main

import (
	"fmt"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

type inputT string

type pointT struct {
	x int
	y int
}

var shape1 = []pointT{{0, 0}, {1, 0}, {2, 0}, {3, 0}}         // -  shape
var shape2 = []pointT{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}} // +  shape
var shape3 = []pointT{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}} // L  shape
var shape4 = []pointT{{0, 0}, {0, 1}, {0, 2}, {0, 3}}         // I  shape
var shape5 = []pointT{{0, 0}, {1, 0}, {0, 1}, {1, 1}}         // [] shape

var shapes = [][]pointT{shape1, shape2, shape3, shape4, shape5}

var WIDTH = 7

func getAbsoluteShape(x0, y0 int, shapePoints []pointT) []pointT {
	shape := []pointT{}

	for _, point := range shapePoints {
		shape = append(shape, pointT{x0 + point.x, y0 + point.y})
	}

	return shape
}

func move(shape []pointT, delta pointT) []pointT {
	movedShape := []pointT{}

	for _, point := range shape {
		movedShape = append(movedShape, pointT{point.x + delta.x, point.y + delta.y})
	}

	return movedShape
}

func moveByJet(shape []pointT, jet string) []pointT {
	delta := pointT{1, 0}

	if jet == "<" {
		delta.x = -1
	}

	return move(shape, delta)
}

func moveDown(shape []pointT) []pointT {
	delta := pointT{0, -1}
	return move(shape, delta)
}

func isBlocked(shape []pointT, blocked map[pointT]bool) bool {
	for _, p := range shape {
		if blocked[p] || p.y == -1 || p.x == -1 || p.x == 7 {
			return true
		}
	}
	return false
}

func getHeight(blocked map[pointT]bool) int {
	max := 0

	for p := range blocked {
		max = utils.Max(p.y, max)
	}

	return max
}

func simulate(jets inputT, nrocks int) int {
	blocked := map[pointT]bool{}

	height := 0
	x0 := 2
	jetIndex := -1

	for n := 0; n < nrocks; n++ {
		y0 := height + 3
		shapePoints := shapes[n%len(shapes)]
		shape := getAbsoluteShape(x0, y0, shapePoints)

		for {
			jetIndex++
			jetIndex = jetIndex % len(jets)
			jet := jets[jetIndex]

			next := moveByJet(shape, string(jet))
			// draw(-1, -1, 9, height+5, shape, blocked)

			if !isBlocked(next, blocked) {
				shape = next
			}

			next = moveDown(shape)
			// draw(-1, -1, 9, height+5, shape, blocked)

			if isBlocked(next, blocked) {
				for _, p := range shape {
					blocked[p] = true
				}
				height = getHeight(blocked) + 1
				break
			}

			shape = next
		}
	}

	return height
}

func draw(x0, y0, w, h int, shape []pointT, blocked map[pointT]bool) {
	for y := utils.Max(y0+h, 10); y >= -1; y-- {
		for x := x0; x < w+x0; x++ {
			ch := "."

			if x == -1 || x == 7 {
				ch = "|"
			}

			if y == -1 {
				ch = "-"
			}

			if blocked[pointT{x, y}] {
				ch = "#"
			}

			for _, p := range shape {
				px := pointT{x, y}
				if p == px {
					ch = "@"
				}
			}

			fmt.Printf("%v", ch)
		}
		fmt.Println("")
	}
}

func part1(input inputT) {
	answer := simulate(input, 2022)
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
