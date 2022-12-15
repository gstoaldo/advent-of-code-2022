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

	re := regexp.MustCompile(`[-\d]+`)

	input := inputT{}

	for _, line := range lines {
		match := re.FindAllString(line, -1)

		xSensor, _ := strconv.Atoi(match[0])
		ySensor, _ := strconv.Atoi(match[1])
		xBeacon, _ := strconv.Atoi(match[2])
		yBeacon, _ := strconv.Atoi(match[3])

		input = append(input, []pointT{{xSensor, ySensor}, {xBeacon, yBeacon}})
	}

	return input
}

func parseFile(path string) inputT {
	return parser(path)
}
