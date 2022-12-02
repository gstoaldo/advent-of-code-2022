package main

import "fmt"

var mapColumn = map[string]string{
	"X": "A",
	"Y": "B",
	"Z": "C",
}

var mapShapeToScore = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

func getRoundScore(opShape string, myShape string) int {
	diff := int(myShape[0]) - int(opShape[0])

	if diff == 1 || diff == -2 {
		return 6
	}

	if diff == 0 {
		return 3
	}

	return 0
}

func getTotalScore(rounds [][]string) int {
	totalScore := 0

	for _, round := range rounds {
		myShape := mapColumn[round[1]]
		shapeScore := mapShapeToScore[myShape]
		roundScore := getRoundScore(round[0], myShape)
		totalScore += shapeScore + roundScore
	}

	return totalScore
}

func getShapeToResult(opShape string, result string) string {
	if result == "Y" {
		return opShape
	}

	if result == "Z" {
		s := string(opShape[0] + 1)

		if s > "C" {
			return "A"
		}

		return s
	}

	s := string(opShape[0] - 1)

	if s < "A" {
		return "C"
	}

	return s
}

func getTotalScoreP2(rounds [][]string) int {
	totalScore := 0

	for _, round := range rounds {
		myShape := getShapeToResult(round[0], round[1])
		shapeScore := mapShapeToScore[myShape]
		roundScore := getRoundScore(round[0], myShape)
		totalScore += shapeScore + roundScore
	}

	return totalScore
}

func part1(input [][]string) {
	answer := getTotalScore(input)
	fmt.Println("part 1:", answer)
}

func part2(input [][]string) {
	answer := getTotalScoreP2(input)
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
