package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
<<<<<<< HEAD
	"time"
=======
>>>>>>> refs/remotes/origin/main
)

// Main Function that will read in answers and compare them to values in CSV file
func main() {
	csvFilename := flag.String("csv", "questions.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	// Use the OS Package to attempt to open file in same directory and error if file cannot be found
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	// Iterare through CSV file and parse lines as questions to be stored in 'questions' slice
	questions := parseLines(lines)

<<<<<<< HEAD
	// Create timer for quiz
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

=======
>>>>>>> refs/remotes/origin/main
	// Declare correct slice, iterate through questions and read answer based on user input
	correct := 0
	for i, qu := range questions {
		fmt.Printf("Question #%d: %s = \n", i+1, qu.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(questions))
			return
		case answer := <-answerCh:
			if answer == qu.a {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(questions))
}

// Function to parse questions
func parseLines(lines [][]string) []question {
	ret := make([]question, len(lines))
	for i, line := range lines {
		ret[i] = question{
			q: line[0],
			a: strings.TrimSpace(line[1]),
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
