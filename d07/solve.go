package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func findDirectoryEndIndex(output inputType, startIndex int) int {
	level := 0

	re := regexp.MustCompile(`\$ cd [\w\/]+`)

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

	re := regexp.MustCompile(`(\d{1,})\s([a-z.]+)`)

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

	re := regexp.MustCompile(`\$ cd [\w\/]+`)

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

func part1(input inputType) {
	answer := getCandidatesSum(input)
	fmt.Println("part 1:", answer)
}

func part2(input inputType) {
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
