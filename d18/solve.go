package main

import (
	"fmt"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

type inputT []cubeT

type cubeT struct {
	x, y, z int
}

func filterByZ(cubes []cubeT, z int) []cubeT {
	filtered := []cubeT{}

	for _, cube := range cubes {
		if cube.z == z {
			filtered = append(filtered, cube)
		}
	}

	return filtered
}

func isTouchingXY(cube1, cube2 cubeT) bool {
	horizontalTouch := (cube1.x == cube2.x) && utils.Abs(cube1.y-cube2.y) == 1
	verticalTouch := (cube1.y == cube2.y) && utils.Abs(cube1.x-cube2.x) == 1

	return horizontalTouch || verticalTouch
}

func isTouchingZ(cube1, cube2 cubeT) bool {
	return cube1.x == cube2.x && cube1.y == cube2.y && utils.Abs(cube1.z-cube2.z) == 1
}

func countTouchingSidesXY(cubesXY []cubeT) int {
	count := 0
	for i := 0; i < len(cubesXY); i++ {
		for j := i + 1; j < len(cubesXY); j++ {
			if isTouchingXY(cubesXY[i], cubesXY[j]) {
				count++
			}
		}
	}

	return count
}

func countFreeSidesXY(cubesXY []cubeT) int {
	return len(cubesXY)*6 - countTouchingSidesXY(cubesXY)*2
}

func countTouchingSidesZ(cubesZ1, cubesZ2 []cubeT) int {
	count := 0
	for _, c1 := range cubesZ1 {
		for _, c2 := range cubesZ2 {
			if isTouchingZ(c1, c2) {
				count++
			}
		}
	}

	return count
}

func countFreeSidesXYZ(cubes []cubeT) int {
	count := 0

	maxZ := 0
	for _, c := range cubes {
		maxZ = utils.Max(maxZ, c.z)
	}

	for z := 0; z <= maxZ; z++ {
		cubesZ1 := filterByZ(cubes, z)
		cubesZ2 := filterByZ(cubes, z+1)

		count += countFreeSidesXY(cubesZ1) - 2*countTouchingSidesZ(cubesZ1, cubesZ2)
	}

	return count
}

func part1(input inputT) {
	answer := countFreeSidesXYZ(input)
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
