package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "questions.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	questions := parseLines(lines)

	correct := 0
	for i, qu := range questions {
		fmt.Printf("Question #%d: %s = \n", i+1, qu.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == qu.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n, correct, len(questions)")
}

func parseLines(lines [][]string) []question {
	ret := make([]question, len(lines))
	for i, line := range lines {
		ret[i] = question{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type question struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println((msg))
	os.Exit(1)
}
