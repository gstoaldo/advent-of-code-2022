package main

import (
	"fmt"
	"sort"

	"github.com/gstoaldo/advent-of-code-2022/utils"
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

func dist(p1 pointT, p2 pointT) int {
	return utils.Abs(p1.x-p2.x) + utils.Abs(p1.y-p2.y)
}

func getXIntervalAtDistance(sensor pointT, dist int, y int) intervalT {
	delta := dist - utils.Abs(y-sensor.y)
	xLower := sensor.x - delta
	xUpper := sensor.x + delta

	if xLower > xUpper {
		return intervalT{} // no solution
	}

	return intervalT{xLower, xUpper}
}

func overlaps(i1 intervalT, i2 intervalT) (bool, intervalT) {
	maxStart := utils.Max(i1.start, i2.start)
	minEnd := utils.Min(i1.end, i2.end)

	minStart := utils.Min(i1.start, i2.start)
	maxEnd := utils.Max(i1.end, i2.end)

	merged := intervalT{minStart, maxEnd}

	return maxStart <= minEnd, merged
}

func mergeIntervals(intervals []intervalT) []intervalT {
	sortedIntervals := append([]intervalT{}, intervals...)

	sort.Slice(sortedIntervals, func(i, j int) bool {
		return sortedIntervals[i].start < sortedIntervals[j].start
	})

	merged := []intervalT{}

	merged = append(merged, sortedIntervals[0])

	for _, interval := range sortedIntervals[1:] {
		if ok, m := overlaps(interval, merged[len(merged)-1]); ok {
			merged[len(merged)-1] = m
		} else {
			merged = append(merged, interval)
		}
	}

	return merged
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
	freeInstervalsBySensor := getFreeIntervalsBySensor(input, y)
	merged := mergeIntervals(freeInstervalsBySensor)

	sum := 0

	for _, interval := range merged {
		sum += interval.end - interval.start + 1
	}

	// remove beacon. There is only one beacon for a given y coordinate
	sum--

	return sum
}

func findDistressBeacon(input inputT, maxCoord int) pointT {
	for y := 0; y < maxCoord; y++ {
		freeInstervalsBySensor := getFreeIntervalsBySensor(input, y)
		merged := mergeIntervals(freeInstervalsBySensor)

		if len(merged) > 1 {
			return pointT{merged[0].end + 1, y}
		}
	}

	return pointT{}
}

func beaconSignal(p pointT) int {
	return p.x*4000000 + p.y
}

func part1(input inputT) {
	// y := 10
	y := 2000000
	answer := countSpacesWithNoBeaconsAtY(input, y)
	fmt.Println("part 1:", answer)
}

func part2(input inputT) {
	// maxCoord := 20
	maxCoord := 4000000
	distressBeacon := findDistressBeacon(input, maxCoord)
	answer := beaconSignal(distressBeacon)
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
