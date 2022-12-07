package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var movesInPattern = `\$ cd [\w\/]+`
var filePattern = `(\d{1,})\s([a-z.]+)`

func findDirectoryEndIndex(output inputType, startIndex int) int {
	level := 0

	re := regexp.MustCompile(movesInPattern)

	for i, line := range output[startIndex:] {
		if re.MatchString(line) {
			level++
			continue
		}

		if line == "$ cd .." {
			level--

			if level == 0 {
				return startIndex + i
			}
		}
	}

	return len(output) - 1
}

func getDirectoryTotalSize(output inputType, startIndex int, endIndex int) int {
	totalSize := 0
	window := output[startIndex : endIndex+1]

	re := regexp.MustCompile(filePattern)

	for _, line := range window {
		match := re.FindStringSubmatch(line)

		if len(match) > 0 {
			fileSize, _ := strconv.Atoi(match[1])

			totalSize += fileSize
		}
	}

	return totalSize
}

func getAllDirectorySize(output inputType) []int {
	sizes := []int{}

	re := regexp.MustCompile(movesInPattern)

	for startIndex, line := range output {
		if re.MatchString(line) {
			endIndex := findDirectoryEndIndex(output, startIndex)
			size := getDirectoryTotalSize(output, startIndex, endIndex)
			sizes = append(sizes, size)
		}
	}

	return sizes
}

func getCandidatesSum(output inputType) int {
	candidatesSum := 0

	sizes := getAllDirectorySize(output)

	for _, size := range sizes {
		if size < 100000 {
			candidatesSum += size
		}
	}

	return candidatesSum
}

func getSizeToFree(output inputType) int {
	spaceAvailable := 70000000
	spaceNecessary := 30000000

	usedSize := getDirectoryTotalSize(output, 0, len(output)-1)
	unusedSize := spaceAvailable - usedSize
	sizeToFree := spaceNecessary - unusedSize

	return sizeToFree
}

func getSmallestDirectorySizeToFree(output inputType) int {
	sizeToFree := getSizeToFree(output)

	sizes := getAllDirectorySize(output)

	minSize := sizes[0] // root directory

	for _, size := range sizes {
		if size >= sizeToFree && size < minSize {
			minSize = size
		}
	}

	return minSize
}

func part1(input inputType) {
	answer := getCandidatesSum(input)
	fmt.Println("part 1:", answer)
}

func part2(input inputType) {
	answer := getSmallestDirectorySizeToFree(input)
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
