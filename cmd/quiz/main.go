package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "", "a csv file in the format of 'question,answer'")
	limit := flag.String("limit", "30", "the time limit for the quiz in seconds")
	flag.Parse()
	if *csvFilename == "" {
		exit("Please provide a csv file")
	}

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to read csv file")
	}
	problems := parseLines(lines)

	d, _ := time.ParseDuration(*limit + "s")
	timer := time.NewTimer(d)
	//TODO Подумать, как можно сделать таймер на каждый вопрос
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\ntimer expired")
			fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("\nYou got %d correct out of %d\n", correct, len(problems))
}

type Problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []Problem {
	res := make([]Problem, len(lines))

	for i, line := range lines {
		res[i] = Problem{
			q: line[0],
			a: line[1],
		}
	}
	return res
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
