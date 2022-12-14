package main

import (
	"io/ioutil"
	"strings"
)

func parser(path string) inputType {
	file, _ := ioutil.ReadFile(path)

	pairs := strings.Split(string(file), "\n\n")

	input := inputType{}
	for _, pair := range pairs {
		input = append(input, strings.Split(pair, "\n"))
	}

	return input
}

func parseFile(path string) inputType {
	return parser(path)
}
