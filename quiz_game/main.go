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

func getNextQuestion(r *csv.Reader) (string, string, bool) {
	problem, err := r.Read()
	if err == io.EOF {
		// end game and show score
		return "", "", true
	}
	if err != nil {
		log.Fatal(err)
	}
	question := problem[0]
	answer := problem[1]
	return question, answer, false
}

func printQuestion(question string) {
	fmt.Printf("question: %s \n", question)
	fmt.Println("answer: ")
}

func manageScore(guess string, answer string, correctGuesses int, incorrectGuess int) (int, int) {
	if guess == answer {
		correctGuesses++
	} else {
		incorrectGuess++
	}
	return correctGuesses, incorrectGuess
}

func printScore(correctGuesses int, incorrectGuess int) {
	fmt.Printf("correct: %d, incorrect: %d \n\n", correctGuesses, incorrectGuess)
}

func StopAndSummariseGame(correctGuesses int, incorrectGuess int) {
	printScore(correctGuesses, incorrectGuess)
	os.Exit(0)
}

func askQuestions() {
	var correctGuesses = 0
	var incorrectGuess = 0
	var start string
	var guess string
	r := readFile()

	fmt.Printf("Hi there, press enter to start the game: \n")
	fmt.Scanln(&start)

	for {
		question, answer, done := getNextQuestion(r)
		if done {
			StopAndSummariseGame(correctGuesses, incorrectGuess)
		}
		printQuestion(question)

		t := time.AfterFunc(3*time.Second, func() {
			StopAndSummariseGame(correctGuesses, incorrectGuess)
		})
		fmt.Scanln(&guess)
		t.Reset(3 * time.Second)
		correctGuesses, incorrectGuess = manageScore(guess, answer, correctGuesses, incorrectGuess)
		fmt.Println()
	}
}

func main() {
	askQuestions()
}
