package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

//Func for exit message
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// struct to hold problems in csv file
type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), // trimming extra spaces if any in problems.csv
		}
	}
	return ret
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "csv file with questions and answers")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the csv file")
	}
	problems := parseLines(lines)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem#%d: %s\n", i+1, p.q) // Printing each questions from csv
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d", correct, len(problems))
}
