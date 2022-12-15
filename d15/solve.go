package main

import (
	"fmt"
)

type inputT [][]pointT

type pointT struct {
	x int
	y int
}

type intervalT struct {
	start int
	end   int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(p1 pointT, p2 pointT) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func getXIntervalAtDistance(sensor pointT, dist int, y int) intervalT {
	delta := dist - abs(y-sensor.y)
	xLower := sensor.x - delta
	xUpper := sensor.x + delta

	if xLower > xUpper {
		return intervalT{} // no solution
	}

	return intervalT{xLower, xUpper}
}

func getFreeIntervalsBySensor(input inputT, y int) []intervalT {
	intervals := []intervalT{}

	for _, line := range input {
		sensor := line[0]
		beacon := line[1]

		interval := getXIntervalAtDistance(sensor, dist(sensor, beacon), y)
		intervals = append(intervals, interval)
	}

	return intervals
}

func countSpacesWithNoBeaconsAtY(input inputT, y int) int {
	countSet := map[int]bool{}

	freeInstervalsBySensor := getFreeIntervalsBySensor(input, y)

	for _, interval := range freeInstervalsBySensor {
		for x := interval.start; x <= interval.end; x++ {
			countSet[x] = true
		}
	}

	// remove beacons from input
	for _, line := range input {
		beacon := line[1]

		_, ok := countSet[beacon.x]

		if ok && beacon.y == y {
			delete(countSet, beacon.x)
		}
	}

	return len(countSet)
}

func part1(input inputT) {
	// y := 10
	y := 2000000
	answer := countSpacesWithNoBeaconsAtY(input, y)
	fmt.Println("part 1:", answer)
}

func part2(input inputT) {
	answer := ""
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("example1.txt")

	part1(input)
	part2(input)
}
