package main

import (
	"io/ioutil"
)

func parser(path string) inputT {
	file, _ := ioutil.ReadFile(path)

	return inputT(string(file))

}

func parseFile(path string) inputT {
	return parser(path)
}
