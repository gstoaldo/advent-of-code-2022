package main

import (
	"io/ioutil"
)

type inputType string

func parser(path string) inputType {
	file, _ := ioutil.ReadFile(path)

	return inputType(string(file))
}

func parseFile(path string) inputType {
	return parser(path)
}
