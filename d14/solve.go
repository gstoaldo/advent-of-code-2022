package main

import (
	"fmt"
	"math"
)

type inputT [][]pointT

type pointT struct {
	x int
	y int
}

type blockedT map[pointT]bool // set

func add(p1 pointT, p2 pointT) pointT {
	return pointT{p1.x + p2.x, p1.y + p2.y}
}

func getRocks(input inputT) blockedT {
	blocked := blockedT{}

	for _, line := range input {
		for i := 1; i < len(line); i++ {
			curr := line[i-1]
			end := line[i]

			blocked[curr] = true

			for curr != end {
				delta := pointT{end.x - curr.x, end.y - curr.y}

				dxLength := int(math.Abs(float64(delta.x)))
				if delta.x == 0 {
					dxLength = 1
				}

				dyLength := int(math.Abs(float64(delta.y)))
				if delta.y == 0 {
					dyLength = 1
				}

				deltaUnit := pointT{delta.x / dxLength, delta.y / dyLength}
				curr = add(curr, deltaUnit)

				blocked[curr] = true
			}
		}
	}

	return blocked
}

func getNextPoint(p0 pointT, rocks blockedT, sands blockedT) pointT {
	for _, delta := range []pointT{{0, 1}, {-1, 1}, {1, 1}} {
		p1 := add(p0, delta)

		_, okRocks := rocks[p1]
		_, okSands := sands[p1]

		if !okRocks && !okSands {
			return p1
		}
	}

	return p0
}

func getMaxY(rocks blockedT) int {
	maxY := 0

	for p, _ := range rocks {
		if p.y > maxY {
			maxY = p.y
		}
	}
	return maxY
}

func simulate(rocks blockedT, source pointT, maxY int) int {
	sands := blockedT{}

	p0 := source

	for p0.y <= maxY {
		p1 := getNextPoint(p0, rocks, sands)

		if p0 == p1 {
			sands[p0] = true
			p0 = source
		} else {
			// draw(494, 0, 20, 20, rocks, sands, p0)
			p0 = p1
		}
	}

	return len(sands)
}

//

func draw(x0 int, y0 int, width int, height int, rocks blockedT, sands blockedT, sand pointT) {
	for y := y0; y < height+y0; y++ {
		for x := x0; x < width+x0; x++ {
			s := "."

			if _, ok := rocks[pointT{x, y}]; ok {
				s = "#"
			}

			if _, ok := sands[pointT{x, y}]; ok {
				s = "o"
			}

			a := pointT{x, y}
			if a == sand {
				s = "o"
			}

			fmt.Printf("%v", s)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func part1(input inputT) {
	rocks := getRocks(input)
	maxY := getMaxY(rocks)
	answer := simulate(rocks, pointT{500, 0}, maxY)
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
