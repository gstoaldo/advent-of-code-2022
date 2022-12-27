package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type inputT []string

func charToDigit(char string) int {
	value, _ := strconv.Atoi(char)

	switch char {
	case "-":
		return -1
	case "=":
		return -2
	default:
		return value
	}
}

func digitToChar(digit int) string {
	switch digit {
	case -2:
		return "="
	case -1:
		return "-"
	default:
		return strconv.Itoa(digit)
	}
}

func toDecimal(snafu string) int {
	sum := 0

	for i, char := range snafu {
		place := len([]rune(snafu)) - i - 1
		value := charToDigit(string(char))
		sum += value * int(math.Pow(5, float64(place)))
	}

	return sum
}

func toSnafu(decimal int) string {
	max := 20
	snafu := make([]string, max)
	carry := 0

	remainder := decimal

	place := max - 1

	for place >= 0 {
		placeValue := int(math.Pow(5, float64(place)))
		quotient := (remainder + carry) / placeValue

		if quotient > 2 || quotient < -2 {
			place++
			carry = int(math.Pow(5, float64(place)))

			if quotient < 0 {
				carry *= -1
			}

			remainder += int(math.Pow(5, float64(place))) * toDecimal(snafu[max-place-1])
			continue
		}

		remainder = remainder - quotient*placeValue
		carry = 0

		snafu[max-place-1] = digitToChar(quotient)
		place--
	}

	snafuStr := fmt.Sprintf("%v", strings.Join(snafu, ""))

	re := regexp.MustCompile(`^0+`)

	return re.ReplaceAllString(snafuStr, "")
}

func sumSnafu(list []string) string {
	sumDecimal := 0

	for _, snafu := range list {
		sumDecimal += toDecimal(snafu)
	}

	return toSnafu(sumDecimal)
}

func part1(input inputT) {
	answer := sumSnafu(input)
	fmt.Println("part 1:", answer)
}

func part2(input inputT) {
	answer := ""
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
