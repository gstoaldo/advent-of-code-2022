package main

import (
	"io/ioutil"
	"strings"
)

type inputType []string

func parser(path string) inputType {
	file, _ := ioutil.ReadFile(path)

	return inputType(strings.Split(string(file), "\n"))
}

func parseFile(path string) inputType {
	return parser(path)
}
