package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "", "a csv file in the format of 'question,answer'")
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

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if strings.TrimSpace(answer) == strings.TrimSpace(p.a) {
			correct++
		}
	}

	fmt.Printf("You got %d correct out of %d\n", correct, len(problems))
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
