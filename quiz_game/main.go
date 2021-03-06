package main

// start game
// did user provide flag for own file to parse? else use problems.csv in same dir as program
// loadQuiz
// present first question to user
// compare anwser ? correct / incorrect
// present next question
// is last question? end game

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var filename string

func init() {
	flag.StringVar(&filename, "filename", "./problems.csv", "define relative path to custom csv questionare file")
	flag.Parse()
}

func readFile() *csv.Reader {
	// set args for examples sake
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)

	return r
}
func askQuestions() {
	// read one row at a time
	r := readFile()
	var correctQuesses = 0
	var incorrectQuess = 0

	for {
		problem, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		question := problem[0]
		answer := problem[1]

		var quess string
		// Taking input from user
		fmt.Printf("question: %s \n", question)
		fmt.Println("answer: ")
		fmt.Scanln(&quess)
		if quess == answer {
			correctQuesses++
		} else {
			incorrectQuess++
		}
		fmt.Println()
	}
	fmt.Printf("correct: %d, incorrect: %d \n\n", correctQuesses, incorrectQuess)
}

func main() {
	askQuestions()
}
