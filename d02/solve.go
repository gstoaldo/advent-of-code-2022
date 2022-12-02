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
	myShapeConverted := mapColumn[myShape]

	diff := int(myShapeConverted[0]) - int(opShape[0])

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
		shapeScore := mapShapeToScore[mapColumn[round[1]]]
		roundScore := getRoundScore(round[0], round[1])
		totalScore += shapeScore + roundScore
	}

	return totalScore
}

func part1(input [][]string) {
	answer := getTotalScore(input)
	fmt.Println("part 1:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
}
