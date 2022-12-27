package main

import (
	"io/ioutil"
	"strings"
)

func parseFile(path string) inputT {
	file, _ := ioutil.ReadFile(path)
	return inputT(strings.Split(string(file), "\n"))
}
