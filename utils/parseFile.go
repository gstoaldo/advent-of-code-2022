package utils

import (
	"bufio"
	"log"
	"os"
)

func ParseFile(path string, parser func(*bufio.Scanner) interface{}) interface{} {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	return parser(scanner)
}
