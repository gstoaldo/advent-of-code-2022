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

	re := regexp.MustCompile(`[A-Z]{2}|\d+`)

	input := inputT{}

	for _, line := range lines {
		match := re.FindAllString(line, -1)

		node, flowStr, linked := match[0], match[1], match[2:]
		flow, _ := strconv.Atoi(flowStr)
		input[node] = nodeT{flow, linked}
	}

	return input
}

func parseFile(path string) inputT {
	return parser(path)
}
