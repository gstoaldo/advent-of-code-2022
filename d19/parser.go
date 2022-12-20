package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func parser(path string) inputT {
	file, _ := ioutil.ReadFile(path)

	lines := strings.Split(string(file), "\n")

	re := regexp.MustCompile(`\d+`)

	bps := inputT{}

	for _, line := range lines {
		mStr := re.FindAllString(line, -1)
		mInt := []int{}

		for _, valStr := range mStr {
			valInt, _ := strconv.Atoi(valStr)
			mInt = append(mInt, valInt)
		}

		r1Ore, r2Ore, r3Ore, r3Clay, r4Ore, r4Obs := mInt[1], mInt[2], mInt[3], mInt[4], mInt[5], mInt[6]

		bp := [][]int{
			{r1Ore, 0, 0, 0},
			{r2Ore, 0, 0, 0},
			{r3Ore, r3Clay, 0, 0},
			{r4Ore, 0, r4Obs, 0},
		}

		bps = append(bps, bp)
	}

	return bps
}

func parseFile(path string) inputT {
	return parser(path)
}
