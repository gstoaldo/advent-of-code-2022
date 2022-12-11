package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var numbersRe = regexp.MustCompile(`\d+`)               // match numbers
var operationRe = regexp.MustCompile(`([\*+])\s(\w+)$`) // match operation symbol and last word

func parser(path string) inputType {
	mmc := 1

	file, _ := ioutil.ReadFile(path)

	chunks := strings.Split(string(file), "\n\n")

	monkeys := []monkey{}

	for _, chunk := range chunks {
		lines := strings.Split(chunk, "\n")

		// starting items
		itemsStr := numbersRe.FindAllString(lines[1], -1)
		items := []int{}

		for _, s := range itemsStr {
			value, _ := strconv.Atoi(s)
			items = append(items, value)
		}

		// operation function
		match := operationRe.FindStringSubmatch(lines[2])
		isOld := match[2] == "old"

		var opsFunc func(old int) int

		if match[1] == "*" {
			opsFunc = func(old int) int {
				v := old
				if !isOld {
					v, _ = strconv.Atoi(match[2])
				}
				return old * v
			}
		} else {
			opsFunc = func(old int) int {
				v := old
				if !isOld {
					v, _ = strconv.Atoi(match[2])
				}
				return (old + v)
			}
		}

		// next func
		divisorStr := numbersRe.FindString(lines[3])
		divisor, _ := strconv.Atoi(divisorStr)

		nextIdxIfTrueStr := numbersRe.FindString(lines[4])
		nextIdxIfFalseStr := numbersRe.FindString(lines[5])

		nextIdxIfTrue, _ := strconv.Atoi(nextIdxIfTrueStr)
		nextIdxIfFalse, _ := strconv.Atoi(nextIdxIfFalseStr)

		testFunc := func(v int) int {
			if v%divisor == 0 {
				return nextIdxIfTrue
			}
			return nextIdxIfFalse
		}

		mmc *= divisor
		monkeys = append(monkeys, monkey{items, opsFunc, testFunc, 0})
	}

	return inputType{monkeys, mmc}
}

func parseFile(path string) inputType {
	return parser(path)
}
