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

func countFreeSides3D(cubes []cubeT) int {
	/*
		Start counting free sides layer by layer (same Z) and subtract matching
		cubes in the next layer (Z+1).
	*/
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

func toSlice(cubes map[cubeT]bool) []cubeT {
	cubesSlice := []cubeT{}

	for c := range cubes {
		cubesSlice = append(cubesSlice, c)
	}

	return cubesSlice
}

func getNeighbours3D(cube cubeT) []cubeT {
	neighbours := []cubeT{}

	for _, delta := range [][]int{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}} {
		neighbours = append(neighbours, cubeT{cube.x + delta[0], cube.y + delta[1], cube.z + delta[2]})
	}

	return neighbours
}

func getAirPocket3D(cubesSet map[cubeT]bool, maxX, maxY, maxZ int, seed cubeT) map[cubeT]bool {
	/*
		Use BFS and start expanding (exploring) from seed. If a point outside the
		limits is reached, that means it is not an air pocket.

		air pocket		not air pocket
										.
		####					##.#
		#..#					#..#
		####					####
	*/
	queue := []cubeT{seed}
	visited := map[cubeT]bool{seed: true}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		neighbours := getNeighbours3D(head)
		for _, n := range neighbours {
			if cubesSet[n] || visited[n] {
				continue
			}

			if n.x < 0 || n.x > maxX || n.y < 0 || n.y > maxY || n.z < 0 || n.z > maxZ {
				return map[cubeT]bool{}
			}

			visited[n] = true
			queue = append(queue, n)
		}
	}

	return visited
}

func getAllAirPockets3D(cubes []cubeT) [][]cubeT {
	maxX, maxY, maxZ := 0, 0, 0
	cubesSet := map[cubeT]bool{}
	allAirPockets := [][]cubeT{}

	for _, c := range cubes {
		maxX = utils.Max(maxX, c.x)
		maxY = utils.Max(maxY, c.y)
		maxZ = utils.Max(maxZ, c.z)

		cubesSet[c] = true
	}

	seeds := map[cubeT]bool{}
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			for z := 0; z <= maxZ; z++ {
				seed := cubeT{x, y, z}
				if !cubesSet[seed] {
					seeds[seed] = true
				}
			}
		}
	}

	for seed := range seeds {
		airPocket := getAirPocket3D(cubesSet, maxX, maxY, maxZ, seed)

		for s := range airPocket {
			delete(seeds, s)
		}

		if len(airPocket) > 0 {
			allAirPockets = append(allAirPockets, toSlice(airPocket))
		}
	}

	return allAirPockets
}

func part1(input inputT) {
	answer := countFreeSides3D(input)
	fmt.Println("part 1:", answer)
}

func part2(input inputT) {
	answer := countFreeSides3D(input)

	airPockets := getAllAirPockets3D(input)
	for _, airPockets := range airPockets {
		answer -= countFreeSides3D(airPockets)
	}

	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
