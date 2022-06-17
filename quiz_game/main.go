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
	"time"
)

var filename string
var timelimit int

func init() {
	flag.StringVar(&filename, "filename", "./problems.csv", "define relative path to custom csv questionare file")
	flag.IntVar(&timelimit, "timelimit", 30000, "define time limit")
	flag.Parse()
}

func readFile() *csv.Reader {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)

	return r
}

func getNextQuestion(r *csv.Reader) (string, string, error) {
	problem, err := r.Read()
	if err == io.EOF {
		return "", "", err
	}
	if err != nil {
		log.Fatal(err)
	}
	question := problem[0]
	answer := problem[1]
	return question, answer, err
}

func askQuestions() {
	// read one row at a time
	var correctQuesses = 0
	var incorrectQuess = 0
	var start string
	var quess string
	r := readFile()

	fmt.Printf("Hi there, press enter to start the game: \n")
	fmt.Scanln(&start)
	for {
		question, answer, err := getNextQuestion(r)
		if err != nil {
			os.Exit(0)
		}
		fmt.Printf("question: %s \n", question)
		fmt.Println("answer: ")

		gameTimer := time.NewTimer(3000 * time.Millisecond)
		go func() {
			println("go func")
			<-gameTimer.C
			gameTimer.Stop()
			fmt.Println("Time expired")
			fmt.Printf("correct: %d, incorrect: %d \n\n", correctQuesses, incorrectQuess)
			os.Exit(0)
		}()
		fmt.Scanln(&quess)
		gameTimer.Reset(3000 * time.Millisecond)
		if quess == answer {
			correctQuesses++
		} else {
			incorrectQuess++
		}
		fmt.Println()
	}
}

func main() {
	askQuestions()
}
