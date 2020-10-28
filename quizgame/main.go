package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {
	records := ParseCSV()
	answersCorrect := 0
	lines := ParseLines(records)
	for i := 0; i < len(lines); i++ {
		s := CheckAnswer(i, lines)
		answersCorrect += s
	}
	answersIncorrect := len(records) - answersCorrect
	fmt.Printf("Number of Correct Answers : %v\n", answersCorrect)
	fmt.Printf("Number of Incorrect Answers: %v\n", answersIncorrect)

}

func ParseCSV() [][]string {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()
	f, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvr := csv.NewReader(f)
	records, err := csvr.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records

}

func ParseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func CheckAnswer(i int, records []problem) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%v = ", records[i].q)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if text == records[i].a {
		return 1
	} else {
		return 0
	}

}
