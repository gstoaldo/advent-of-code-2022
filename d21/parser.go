package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var reNode = regexp.MustCompile(`(\w+): (\w+) ([+-\/\*]) (\w+)`)
var reValue = regexp.MustCompile(`(\w+): (\d+)`)

func createNode(target string, lines []string) nodeT {
	line := ""
	for _, l := range lines {
		if l[:4] == target {
			line = l
			break
		}
	}

	valueMatch := reValue.FindStringSubmatch(line)

	if len(valueMatch) > 0 {
		name, valueStr := valueMatch[1], valueMatch[2]
		value, _ := strconv.Atoi(valueStr)

		return nodeT{name: name, value: value}
	}

	nodeMatch := reNode.FindStringSubmatch(line)

	name, left, operator, right := nodeMatch[1], nodeMatch[2], nodeMatch[3], nodeMatch[4]

	leftNode := createNode(left, lines)
	rightNode := createNode(right, lines)

	return nodeT{
		name:     name,
		operator: operator,
		value:    0,
		left:     &leftNode,
		right:    &rightNode,
	}
}

func parser(path string) inputT {
	file, _ := ioutil.ReadFile(path)

	lines := strings.Split(string(file), "\n")

	return inputT(createNode("root", lines))
}

func parseFile(path string) inputT {
	return parser(path)
}
