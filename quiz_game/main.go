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
	"fmt"
	"io"
	"log"
	"os"
)

func readFile() *csv.Reader {
	// set args for examples sake
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)

	return r
}
func askQuestions() {
	// read one row at a time
	r := readFile()
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
		fmt.Printf("what is %s \n", question)
		fmt.Println("answer: ")
		fmt.Scanln(&quess)
		fmt.Printf("Question: %s -> answer: %v -> quess: %v\n", question, answer, quess)
	}
}

func main() {
	askQuestions()
}
